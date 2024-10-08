#ifndef WRAPPER_GUI_QT_WEBVIEW_HEADER
#define WRAPPER_GUI_QT_WEBVIEW_HEADER

#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif
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
        bool developmentMode
    );
#ifdef __cplusplus
}
#endif

#endif