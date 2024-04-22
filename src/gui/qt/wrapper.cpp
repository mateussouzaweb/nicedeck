#include "wrapper.hpp"
#include "webview.hpp"
#include <stdbool.h>

// Start application
int start_qt_app_wrapper (
    const char *appVendor,
    const char *appName,
    const char *appId,
    const char *appIcon,
    const char *appUrl,
    const char *appVersion,
    bool windowFullScreen,
    bool windowMaximized,
    bool windowDecorated,
    int windowWidth,
    int windowHeight,
    bool developMode
){
    return start_qt_app(
        appVendor, 
        appName, 
        appId, 
        appIcon,
        appUrl,
        appVersion,
        windowFullScreen,
        windowMaximized, 
        windowDecorated,
        windowWidth,
        windowHeight,
        developMode
    );
}