#ifndef GUI_GTK_WEBVIEW_HEADER
#define GUI_GTK_WEBVIEW_HEADER

int start_gtk_app (
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
    bool developMode,
    bool showInspector
);

#endif