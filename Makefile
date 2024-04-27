export BACKEND_PORT=4000
export FRONTEND_PORT=3000
export GO_MOD=goapp
export GO_MOD_OLD=mytest

setup:
		go mod edit -module $(GO_MOD)
		find . -type f -name '*.go' -exec sed -i '' -e 's/$(GO_MOD_OLD)/$(GO_MOD)/g' {} +
		go mod tidy
		cd ui && npm install && npm run build
		@echo "SETUP IS DONE!"
run:
		cd ui && npm run dev &
		@echo "Waiting for frontend server to start..."
		@while ! nc -z localhost $(FRONTEND_PORT); do sleep 1; done
		export DEV=true && go mod tidy && go run .
preview:
		cd ui && npm run build
		export DEV=false && go mod tidy && go run .
build:
		cd ui && npm run build
		export DEV=false && go mod tidy && go build
build-mac:
		cd ui && npm run build
		export DEV=false && go mod tidy && GOOS=darwin go build
build-lin:
		cd ui && npm run build
		export DEV=false && go mod tidy && GOOS=linux go build
build-win:
		cd ui && npm run build
		export DEV=false && go mod tidy && GOOS=windows go build
prod-run:
		./$(GO_MOD)
