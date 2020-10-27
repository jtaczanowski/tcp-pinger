.PHONY: build build-windows build-macos build-linux lint lint-deps

build: build-windows build-macos build-linux



build-windows:
	@echo "Building application for GOOS=windows GOARCH=amd64"
	GOOS=windows GOARCH=amd64 go build -o bin/Windows/amd64/tcp-pinger.exe ./cmd/tcp-pinger
	@echo "Building application for GOOS=windows GOARCH=386"
	GOOS=windows GOARCH=amd64 go build -o bin/Windows/386/tcp-pinger.exe ./cmd/tcp-pinger

build-macos:
	@echo "Building application for GOOS=macos GOARCH=amd64"
	GOOS=darwin GOARCH=amd64 go build -o bin/MacOS/amd64/tcp-pinger ./cmd/tcp-pinger

build-linux:
	@echo "Building application for GOOS=linux GOARCH=amd64"
	GOOS=linux GOARCH=amd64 go build -o bin/Linux/amd64/tcp-pinger ./cmd/tcp-pinger
	@echo "Building application for GOOS=linux GOARCH=386"
	GOOS=linux GOARCH=386 go build -o bin/Linux/386/tcp-pinger ./cmd/tcp-pinger
	@echo "Building application for GOOS=linux GOARCH=arm GOARM=5"
	GOOS=linux GOARCH=arm GOARM=5 go build -o bin/Linux/arm5/tcp-pinger ./cmd/tcp-pinger

lint-deps:
		go get github.com/golangci/golangci-lint/cmd/golangci-lint

lint: lint-deps
	golangci-lint run