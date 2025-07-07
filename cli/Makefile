.PHONY: build test clean install dev release lint

# Variáveis
APP_NAME := ltx
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

# Build local
build:
	go build $(LDFLAGS) -o $(APP_NAME) ./main.go

# Build para todas as plataformas
build-all:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o bin/$(APP_NAME)-linux-amd64 ./main.go
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o bin/$(APP_NAME)-linux-arm64 ./main.go
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o bin/$(APP_NAME)-darwin-amd64 ./main.go
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o bin/$(APP_NAME)-darwin-arm64 ./main.go
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o bin/$(APP_NAME)-windows-amd64.exe ./main.go

# Testes
test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Desenvolvimento
dev:
	go run ./main.go

# Linting
lint:
	golangci-lint run

# Instalação local (copia para ../bin/)
install: build
	mkdir -p ../bin
	cp $(APP_NAME) ../bin/

# Limpeza
clean:
	rm -f $(APP_NAME)
	rm -rf bin/
	rm -f coverage.out coverage.html

# Release com GoReleaser
release:
	goreleaser release --clean

# Release snapshot (teste)
release-snapshot:
	goreleaser release --snapshot --clean

# Setup de desenvolvimento
setup-dev:
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/goreleaser/goreleaser@latest

# Formato de código
fmt:
	go fmt ./...

# Verificar módulos
mod-tidy:
	go mod tidy
