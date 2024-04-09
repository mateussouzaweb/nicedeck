LINUX_DEPLOY_SRC = https://github.com/linuxdeploy/linuxdeploy/releases/download/continuous/linuxdeploy-x86_64.AppImage
LINUX_DEPLOY_PLUGIN_GTK_SRC = https://raw.githubusercontent.com/linuxdeploy/linuxdeploy-plugin-gtk/master/linuxdeploy-plugin-gtk.sh

run:
	go run cmd/main.go
	
build:
	mkdir -p bin/
	go build -o bin/nicedeck cmd/main.go

app-image: clean build
	wget $(LINUX_DEPLOY_SRC) -O bin/LinuxDeploy.AppImage
	wget $(LINUX_DEPLOY_PLUGIN_GTK_SRC) -O bin/linuxdeploy-plugin-gtk.sh
	chmod +x bin/LinuxDeploy.AppImage bin/linuxdeploy-plugin-gtk.sh
	./bin/LinuxDeploy.AppImage \
		--appdir bin/AppDir \
		--executable bin/nicedeck \
		--desktop-file src/nicedeck/resources/com.mateussouzaweb.NiceDeck.desktop \
		--icon-file src/nicedeck/resources/nicedeck.svg \
		--plugin gtk \
		--output appimage
	mv ./NiceDeck-*.AppImage bin/

deploy: app-image
	mkdir -p $(HOME)/Applications
	cp bin/NiceDeck-*.AppImage $(HOME)/Applications/NiceDeck.AppImage
	chmod +x $(HOME)/Applications/NiceDeck.AppImage

clean:
	[ -d bin/ ] && rm -r bin/ || true