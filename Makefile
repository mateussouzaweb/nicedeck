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

flatpak-deps:
	sudo apt install -y flatpak flatpak-builder
	flatpak remote-add --if-not-exists flathub https://flathub.org/repo/flathub.flatpakrepo
	flatpak install flathub org.gnome.Platform//46
	flatpak install flathub org.gnome.Sdk//46
	flatpak install flathub org.freedesktop.Sdk.Extension.golang//23.08

flatpak-build:
	flatpak-builder --force-clean --repo=.flatpak-repository .flatpak-build-dir flatpak/manifest.yml

flatpak-bundle:
	mkdir -p bin/
	flatpak build-bundle .flatpak-repository bin/nicedeck.flatpak com.mateussouzaweb.NiceDeck

flatpak-install:
	flatpak-builder --user --install --force-clean .flatpak-build-dir flatpak/manifest.yml
	flatpak run com.mateussouzaweb.NiceDeck