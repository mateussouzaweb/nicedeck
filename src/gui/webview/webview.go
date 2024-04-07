package webview

// Required packages:
// sudo apt install -y libgtk-4-dev libwebkitgtk-6.0-dev
// sudo dnf install -y gtk-4-devel webkitgtk-6.0-devel

// #cgo pkg-config: gtk4 webkitgtk-6.0
// #include "webview.h"
import "C"

// Open webview application by calling it from c++ app
func Open(address string, width int, height int) error {
	C.start_application(0, nil)
	return nil
}
