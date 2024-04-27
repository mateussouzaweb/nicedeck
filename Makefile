deps:
	sudo apt update
	sudo apt install -y libgtk-4-dev libwebkitgtk-6.0-dev
	sudo apt install -y qt6-base-dev qt6-webengine-dev

run:
	go run cmd/main.go

clean:
	[ -d bin/ ] && rm -r bin/ || true

build: clean
	mkdir -p bin/
	go build -o bin/nicedeck cmd/main.go

install: build
	mkdir -p $(HOME)/Applications
	cp bin/nicedeck $(HOME)/Applications/NiceDeck
	chmod +x $(HOME)/Applications/NiceDeck
	$(HOME)/Applications/NiceDeck