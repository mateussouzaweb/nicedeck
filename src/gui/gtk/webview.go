package gtk

// Required packages:
// sudo apt install -y libgtk-4-dev libwebkitgtk-6.0-dev
// sudo dnf install -y gtk4-devel webkitgtk6.0-devel

// #cgo pkg-config: gtk4 webkitgtk-6.0
// #include <stdio.h>
// #include <stdlib.h>
// #include <stdbool.h>
// #include "wrapper.hpp"
import "C"
import (
	"runtime"
	"unsafe"
)

// Open UI as GTK application mode
func Open(url string, version string, developmentMode bool) error {

	runtime.LockOSThread()

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

	defer func() {
		C.free(unsafe.Pointer(appName))
		C.free(unsafe.Pointer(appID))
		C.free(unsafe.Pointer(appIcon))
		C.free(unsafe.Pointer(appURL))
		C.free(unsafe.Pointer(appVersion))
	}()

	C.start_gtk_app_wrapper(
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
		(C.bool)(developmentMode),
	)

	return nil
}
