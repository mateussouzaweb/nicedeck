run:
	go run cmd/main.go
	
build:
	mkdir -p bin/
	go build -o bin/NiceDeck cmd/main.go

deploy: build
	mkdir -p $(HOME)/Applications
	cp bin/NiceDeck $(HOME)/Applications/NiceDeck
	chmod +x $(HOME)/Applications/NiceDeck