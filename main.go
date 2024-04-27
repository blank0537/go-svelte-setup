package main

import (
	"mytest/ui"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var envVar string = os.Getenv("DEV")
var port = func() string {
	if p := os.Getenv("BACKEND_PORT"); p != "" {
		return p
	}
	return "4000"
}()
var clientport string = os.Getenv("FRONTEND_PORT")
var host string = "localhost"
var schema string = "http"

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	if envVar == "true" {
		// Reverse proxy for frontend server
		proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: schema, Host: host + ":" + clientport})
		e.Any("/*", echo.WrapHandler(proxy))
	} else {
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Root:       "/",
			Index:      "index.html",
			Browse:     false,
			HTML5:      true,
			Filesystem: http.FS(ui.DistDirFS),
		}))
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	regular := color.New()
	regular.Printf("\n")
	regular.Printf("├─ BACKEND: %s\n", color.CyanString("%s://%s:%s/ping", schema, host, port))
	if envVar == "true" {
		regular.Printf("├─ FRONTEND: %s\n", color.CyanString("%s://%s:%s", schema, host, clientport))
	}
	regular.Printf("└─ GO + SVELTE: %s\n", color.CyanString("%s://%s:%s", schema, host, port))
	e.Logger.Fatal(e.Start(":" + port))
}
