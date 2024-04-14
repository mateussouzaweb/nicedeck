package webview

// Required packages:
// sudo apt install -y libgtk-4-dev libwebkitgtk-6.0-dev
// sudo dnf install -y gtk-4-devel webkitgtk-6.0-devel

// #cgo pkg-config: gtk4 webkitgtk-6.0
// #include <stdio.h>
// #include <stdlib.h>
// #include "webview.h"
import "C"
import "unsafe"

// Open UI in webview mode
func Open(url string, version string) error {

	appName := C.CString("NiceDeck")
	appID := C.CString("com.mateussouzaweb.NiceDeck")
	appIcon := C.CString("nicedeck")
	appURL := C.CString(url)
	appVersion := C.CString(version)
	windowFullScreen := (C.bool)(false)
	windowMaximized := (C.bool)(true)
	windowDecorated := (C.bool)(true)
	windowWidth := (C.int)(1280)
	windowHeight := (C.int)(800)
	developMode := (C.bool)(true)
	showInspector := (C.bool)(false)

	defer func() {
		C.free(unsafe.Pointer(appName))
		C.free(unsafe.Pointer(appID))
		C.free(unsafe.Pointer(appIcon))
		C.free(unsafe.Pointer(appURL))
		C.free(unsafe.Pointer(appVersion))
	}()

	C.start_app(
		appName,
		appID,
		appIcon,
		appURL,
		appVersion,
		windowFullScreen,
		windowMaximized,
		windowDecorated,
		windowWidth,
		windowHeight,
		developMode,
		showInspector,
	)

	return nil
}
