#include "webview.hpp"
#include <QApplication>
#include <QLoggingCategory>
#include <QMainWindow>
#include <QShortcut>
#include <QWebEngineView>
#include <QWebEngineSettings>

// WebApplication struct
typedef struct {
    QApplication *app;
    QMainWindow *window;
    QWebEngineView *webview;
    QWebEngineSettings *settings;
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
} WebApplication;

// Start application
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

    // Create application
    char* argv[] = {};
    int argc = 0;

    QApplication app(argc, argv);
    instance.app = &app;

    app.setOrganizationName(instance.appVendor);
    app.setApplicationName(instance.appId);
    app.setApplicationDisplayName(instance.appName);
    app.setApplicationVersion(instance.appVersion);
    app.setDesktopFileName(instance.appId);
    app.setWindowIcon(QIcon::fromTheme(instance.appIcon));

    // Create main window
    QMainWindow window;
    instance.window = &window;

    window.setWindowTitle(instance.appName);
    window.setWindowIcon(QIcon::fromTheme(appIcon));
    window.setWindowIconText(instance.appIcon);
    window.setWindowFlag(Qt::FramelessWindowHint, !instance.windowDecorated);
    window.setMinimumSize(instance.windowWidth / 2, instance.windowHeight / 2);
    window.resize(instance.windowWidth, instance.windowHeight);

    // Show fullscreen, maximized on in normal mode
    if (instance.windowFullScreen) {
        window.showFullScreen();
    } else if (instance.windowMaximized) {
        window.showMaximized();
    } else {
        window.show();
    }

    // Attack keyboard shortcuts
    QShortcut *shortcut = new QShortcut(QKeySequence(QKeySequence::Quit), &window);
    QObject::connect(shortcut, &QShortcut::activated, &window, &app.quit);

    // Setup logging
    QLoggingCategory contextLog = QLoggingCategory("qt.webenginecontext");
    contextLog.setFilterRules("*.info=false");

    // Enable developer mode
    if (instance.developMode){
        qputenv("QTWEBENGINE_REMOTE_DEBUGGING", "0.0.0.0:9090");
    }

    // Create webview
    QWebEngineView webview(&window);
    instance.webview = &webview;
    window.setCentralWidget(&webview);

    // Set webview settings
    QWebEngineSettings *settings = webview.settings();
    instance.settings = settings;

    settings->setAttribute(settings->JavascriptEnabled, true);
    settings->setAttribute(settings->JavascriptCanAccessClipboard, true);
    settings->setAttribute(settings->LocalStorageEnabled, true);

    // Load target URL
    webview.show();
    webview.load(QUrl(instance.appUrl));

    // Run application
    int status = app.exec();

    return status;
}