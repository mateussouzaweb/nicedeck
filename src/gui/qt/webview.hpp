#ifndef GUI_QT_WEBVIEW_HEADER
#define GUI_QT_WEBVIEW_HEADER

int start_qt_app (
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
);

#endif