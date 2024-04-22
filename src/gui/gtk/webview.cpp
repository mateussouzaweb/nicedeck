#include "webview.hpp"
#include <gtk/gtk.h>
#include <gdk/gdk.h>
#include <webkit/webkit.h>

// WebApplication struct
typedef struct {
    GtkApplication *app;
    GtkWidget *window;
    GtkWidget *webview;
    WebKitSettings *settings;
    WebKitWebInspector *inspector;
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

// On quit activated callback
static void on_quit_activated(GSimpleAction *action, GVariant *parameter, gpointer data)
{
    GtkApplication *app = (GtkApplication*) data;
    g_application_quit(G_APPLICATION(app));
}

// On startup callback
static void on_startup_app(GtkApplication *app, gpointer data)
{
    // App global actions
    static GActionEntry app_entries[] = {{"quit", on_quit_activated, NULL, NULL, NULL}};
    g_action_map_add_action_entries(G_ACTION_MAP(app), app_entries, G_N_ELEMENTS(app_entries), app);

    // App keyboard shortcuts
    const char *quit_accels[2] = {"<Ctrl>Q", NULL};
    gtk_application_set_accels_for_action(GTK_APPLICATION(app), "app.quit", quit_accels);
}

// On activate callback
static void on_activate_app(GtkApplication *app, gpointer data)
{
    WebApplication *instance = (WebApplication*) data;

    // Create window
    instance->window = gtk_application_window_new(app);
    gtk_window_set_title(GTK_WINDOW(instance->window), instance->appName);
    gtk_window_set_icon_name(GTK_WINDOW(instance->window), instance->appIcon);
    gtk_window_set_decorated(GTK_WINDOW(instance->window), instance->windowDecorated);
    gtk_window_set_default_size(GTK_WINDOW(instance->window), instance->windowWidth, instance->windowHeight);
    gtk_widget_set_size_request(GTK_WIDGET(instance->window), instance->windowWidth / 2, instance->windowHeight / 2);
    gtk_window_set_resizable(GTK_WINDOW(instance->window), true);

    // Set fullscreen or maximize window
    if (instance->windowFullScreen) {
        gtk_window_fullscreen(GTK_WINDOW(instance->window));
    } else if (instance->windowMaximized) {
        gtk_window_maximize(GTK_WINDOW(instance->window));
    }

    // Show window
    gtk_window_present(GTK_WINDOW(instance->window));

    // Create webview
    instance->webview = webkit_web_view_new();
    gtk_window_set_child(GTK_WINDOW(instance->window), instance->webview);

    // Set webview settings
    instance->settings = webkit_web_view_get_settings(WEBKIT_WEB_VIEW(instance->webview));
    webkit_settings_set_user_agent_with_application_details(instance->settings, instance->appId, instance->appVersion);
    webkit_settings_set_enable_javascript(instance->settings, true);
    webkit_settings_set_javascript_can_access_clipboard(instance->settings, true);
    webkit_settings_set_enable_html5_local_storage(instance->settings, true);
    webkit_settings_set_enable_write_console_messages_to_stdout(instance->settings, false);
    webkit_settings_set_enable_developer_extras(instance->settings, instance->developMode);

    // Show inspector
    if( instance->developMode ){
        instance->inspector = webkit_web_view_get_inspector(WEBKIT_WEB_VIEW(instance->webview));
        webkit_web_inspector_show(WEBKIT_WEB_INSPECTOR(instance->inspector));
    }

    // Load target URL
    webkit_web_view_load_uri(WEBKIT_WEB_VIEW(instance->webview), instance->appUrl);

}

// Start application
int start_gtk_app(
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
)
{
    // WebApplication struct instance
    static WebApplication instance;

    // Set variables
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

    // Create GTK application
    instance.app = gtk_application_new(appId, G_APPLICATION_DEFAULT_FLAGS);

    // Attach signal callbacks
    g_signal_connect(instance.app, "startup", G_CALLBACK(on_startup_app), &instance);
    g_signal_connect(instance.app, "activate", G_CALLBACK(on_activate_app), &instance);

    // Run GTK application
    int status = g_application_run(G_APPLICATION(instance.app), 0, NULL);
    g_object_unref(instance.app);

    return status;
}