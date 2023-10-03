run:
	go run cmd/main.go
	
build:
	go build -o bin/nicedeck cmd/main.go

deploy: build
	sudo cp bin/nicedeck $(HOME)/.local/bin/nicedeck
	sudo chmod +x $(HOME)/.local/bin/nicedeck