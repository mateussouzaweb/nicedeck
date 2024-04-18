#ifndef GUI_QT_WEBVIEW_HEADER
#define GUI_QT_WEBVIEW_HEADER

#include <QApplication>
#include <QMainWindow>
#include <QWebEngineView>
#include <QWebEngineSettings>
#include <QWebEngineProfile>

// WebApplication struct
typedef struct {
    QApplication *app;
    QMainWindow *window;
    QWebEngineView *webview;
    QWebEngineSettings *settings;
    QWebEngineProfile *profile;
    const char *appVendor;
    const char *appName;
    const char *appId;
    const char *appIcon;
    const char *appUrl;
    const char *appVersion;
    bool windowFullScreen;
    bool windowMaximized;
    bool windowDecorated;
    int windowWidth;
    int windowHeight;
    bool developMode;
    bool showInspector;
} WebApplication;

// Start application
int start_app (
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
    bool developMode,
    bool showInspector
){

    // WebApplication struct instance
    static WebApplication instance;

    // Set variables
    instance.appVendor = appVendor;
    instance.appName = appName;
    instance.appId = appId;
    instance.appIcon = appIcon;
    instance.appUrl = appUrl;
    instance.appVersion = appVersion;
    instance.windowFullScreen = windowFullScreen;
    instance.windowMaximized = windowMaximized;
    instance.windowDecorated = windowDecorated;
    instance.windowWidth = windowWidth;
    instance.windowHeight = windowHeight;
    instance.developMode = developMode;
    instance.showInspector = showInspector;

    std::string title = instance.appName;
    char* argv[] = { title.data() };
    int argc = 1;

    const auto vendor = QString::fromUtf8(instance.appVendor);
    const auto product = QString::fromUtf8(instance.appId);
    const auto version = QString::fromUtf8(instance.appVersion);
    const auto icon = QString::fromUtf8(instance.appIcon);

    QCoreApplication::setOrganizationName(vendor);
    QCoreApplication::setApplicationName(product);
    QCoreApplication::setApplicationVersion(version);

    QApplication app(argc, argv);
    app.setWindowIcon(QIcon::fromTheme(icon));

    QWebEngineView webview;
    webview.setMinimumSize(instance.windowWidth / 2, instance.windowHeight / 2);
    webview.resize(instance.windowWidth, instance.windowHeight);

    if (!instance.windowDecorated) {
        webview.setWindowFlag(Qt::FramelessWindowHint, true);
    }

    if (instance.windowFullScreen) {
        webview.setWindowState(Qt::WindowFullScreen);
    } else if (instance.windowMaximized) {
        webview.setWindowState(Qt::WindowMaximized);
    }

    webview.load(QUrl(QString::fromStdString(instance.appUrl)));
    webview.show();

    return app.exec();
}

#endif