LINUX_DEPLOY_SRC = https://github.com/linuxdeploy/linuxdeploy/releases/download/continuous/linuxdeploy-x86_64.AppImage

run:
	go run cmd/main.go
	
build:
	mkdir -p bin/
	go build -o bin/nicedeck cmd/main.go

app-image: clean build
	wget $(LINUX_DEPLOY_SRC) -O bin/LinuxDeploy.AppImage
	chmod +x bin/LinuxDeploy.AppImage
	./bin/LinuxDeploy.AppImage \
		--appdir bin/AppDir \
		--executable bin/nicedeck \
		--desktop-file src/nicedeck/resources/com.mateussouzaweb.NiceDeck.desktop \
		--icon-file src/nicedeck/resources/nicedeck.svg \
		--output appimage
	mv ./NiceDeck-*.AppImage bin/

deploy: app-image
	mkdir -p $(HOME)/Applications
	cp bin/NiceDeck-*.AppImage $(HOME)/Applications/NiceDeck.AppImage
	chmod +x $(HOME)/Applications/NiceDeck.AppImage

clean:
	[ -d bin/ ] && rm -r bin/ || true