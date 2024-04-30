#include "wrapper.hpp"
#include "webview.hpp"
#include <stdbool.h>

// Start application
int start_gtk_app_wrapper (
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
    bool developmentMode
){
    return start_gtk_app(
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
        developmentMode
    );
}