# GO + SVELTE SETUP

Clone this repo for initial Go and Svelte (SPA Routing) setup. It produces single file to run your both Go & Svelte on single port

### Project Includes

-   Go - Backend
-   Svelte - Frontend
-   echo - For go backend apis and serve svelte spa app
-   svelte-routing - npm package for routing in SPA
-   makefile - Various commands to setup, run & build

## Setup & Run

Default Project Configs

-   frontend port - 3000
-   backend port - 4000
-   Go module name - goapp

Open Makefile to change these default values as per your need

-   Env varialbles in makefile
    1. export BACKEND_PORT=4000 -> backend port
    2. export FRONTEND_PORT=3000 -> frontend port
    3. export GO_MOD=goapp -> go module name(this is current name for go mod)
    4. export GO_MOD_OLD=mytest -> go old module name(if you want to edit go mod name then please set previous name here to that it will update moudle in all go files)

Run setup command - makefile setup (If you see an error at ui/embed.go path on "all:dist" then ignore)

Once setup is you can use below command for other operations

-   For local development - makefile run
-   For preview - makefile preview
-   Build Commands
    1. For default build - makefile build
    2. For windows build - makefile win
    3. For linux build - makefile lin
    4. For mac build - makefile mac
-   Once build is complete you can run your binary to check by running - makefile prod-run
