package qt

// Required packages:
// sudo apt install -y qt6-base-dev qt6-webengine-dev
// sudo dnf install -y qt6-qtbase-devel qt6-qtwebengine-devel

// #cgo pkg-config: Qt6Core Qt6Widgets Qt6Network Qt6WebEngineCore Qt6WebEngineWidgets
// #include <stdio.h>
// #include <stdlib.h>
// #include "webview.hpp"
import "C"
import "unsafe"

// Open UI in QT application mode
func Open(url string, version string) error {

	appVendor := C.CString("com.mateussouzaweb")
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
		C.free(unsafe.Pointer(appVendor))
		C.free(unsafe.Pointer(appName))
		C.free(unsafe.Pointer(appID))
		C.free(unsafe.Pointer(appIcon))
		C.free(unsafe.Pointer(appURL))
		C.free(unsafe.Pointer(appVersion))
	}()

	C.start_app(
		appVendor,
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
