ARCH=$(shell go env GOARCH)
SYSTEM=$(shell go env GOOS)

run:
	go run cmd/main.go --dev --gui=headless

clean:
	[ -d bin/ ] && rm -r bin/ || true

build: clean
	mkdir -p bin/

	# Linux
	GOOS=linux GOARCH=amd64 \
	go build -o bin/nicedeck-linux-amd64 cmd/main.go
	GOOS=linux GOARCH=arm64 \
	go build -o bin/nicedeck-linux-arm64 cmd/main.go

	# MacOS
	GOOS=darwin GOARCH=amd64 \
	go build -o bin/nicedeck-macos-amd64 cmd/main.go
	GOOS=darwin GOARCH=arm64 \
	go build -o bin/nicedeck-macos-arm64 cmd/main.go

	# Windows
	GOOS=windows GOARCH=amd64 \
	go build -ldflags="-H windowsgui" -o bin/nicedeck-windows-amd64.exe cmd/main.go
	GOOS=windows GOARCH=arm64 \
	go build -ldflags="-H windowsgui" -o bin/nicedeck-windows-arm64.exe cmd/main.go
