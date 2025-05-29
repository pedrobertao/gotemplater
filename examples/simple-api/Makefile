APP_NAME=Makefile

.PHONY: run build build-windows build-darwin build-linux clean

run:
	go run main.go

build: build-windows build-darwin build-linux

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME).exe

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(APP_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o $(APP_NAME)-darwin-arm64

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)-linux-amd64

clean:
	rm -f $(APP_NAME).exe $(APP_NAME)-darwin-amd64 $(APP_NAME)-darwin-arm64 $(APP_NAME)-linux-amd64

help:
	@echo "Available targets:"
	@echo "  run           - Run the application"
	@echo "  build         - Build for all platforms"
	@echo "  build-windows - Build for Windows"
	@echo "  build-darwin  - Build for macOS (Intel and Apple Silicon)"
	@echo "  build-linux   - Build for Linux"
	@echo "  clean         - Remove all built binaries"
	@echo "  help          - Show this help message"