.PHONY: all format lint clean

GOCMD=go
LDFLAGS="-s -w ${LDFLAGS_OPT}"

all: build format lint ## Format, lint and build

build: ## Build
	go build -o bin/main main.go

compile: ## Compile for every OS and Platform
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/ai-commit-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/ai-commit-darwin-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/ai-commit-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/ai-commit-linux-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/ai-commit-windows-amd64 main.go
	GOOS=windows GOARCH=arm64 go build -o bin/ai-commit-windows-arm64 main.go

format: ## Format code
	${GOCMD} fmt ./...

lint: ## Run linter
	golangci-lint run

clean: ## Cleanup build dir
	rm -r bin/

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
