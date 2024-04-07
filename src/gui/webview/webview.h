#include <gtk/gtk.h>
#include <gdk/gdk.h>
#include <webkit/webkit.h>

// App info
static const char *APP_NAME = "NiceDeck";
static const char *APP_ID = "com.mateussouzaweb.NiceDeck";
static const char *APP_ICON = "nicedeck";
static const char *APP_URL = "http://127.0.0.1:14935";

// Display mode: default, maximized or fullscreen
static const char *APP_DISPLAY_MODE = "maximized"; 
static const bool APP_DECORATED = true;
static const int APP_DEFAULT_WIDTH = 1280;
static const int APP_DEFAULT_HEIGHT = 720;
static const int APP_MIN_WIDTH = 600;
static const int APP_MIN_HEIGHT = 300;

// Quit activated callback
static void quit_activated(GSimpleAction *action, GVariant *parameter, gpointer app)
{
  g_application_quit(G_APPLICATION(app));
}

// On startup callback
static void on_startup_app(GtkApplication *app)
{
  // App global actions
  static GActionEntry app_entries[] = {{"quit", quit_activated, NULL, NULL, NULL}};
  g_action_map_add_action_entries(G_ACTION_MAP(app), app_entries, G_N_ELEMENTS(app_entries), app);

  // App keyboard shortcuts
  const char *quit_accels[2] = {"<Ctrl>Q", NULL};
  gtk_application_set_accels_for_action(GTK_APPLICATION(app), "app.quit", quit_accels);
}

// On activate callback
static void on_activate_app(GtkApplication *app)
{
  // Window
  GtkWidget *window = gtk_application_window_new(app);
  gtk_window_set_application(GTK_WINDOW(window), GTK_APPLICATION(app));
  gtk_window_set_title(GTK_WINDOW(window), APP_NAME);
  gtk_window_set_icon_name(GTK_WINDOW(window), APP_ICON);

  gtk_window_set_decorated(GTK_WINDOW(window), APP_DECORATED);
  gtk_window_set_default_size(GTK_WINDOW(window), APP_DEFAULT_WIDTH, APP_DEFAULT_HEIGHT);
  gtk_widget_set_size_request(GTK_WIDGET(window), APP_MIN_WIDTH, APP_MIN_HEIGHT);
  gtk_window_set_resizable(GTK_WINDOW(window), true);
  gtk_application_window_set_show_menubar(GTK_APPLICATION_WINDOW(window), true);

  // Create webview
  GtkWidget *webview = webkit_web_view_new();
  gtk_window_set_child(GTK_WINDOW(window), webview);

  // Set webview settings
  WebKitSettings *settings = webkit_web_view_get_settings(WEBKIT_WEB_VIEW(webview));
  webkit_settings_set_user_agent_with_application_details(settings, APP_ID, "");
  webkit_settings_set_enable_javascript(settings, true);
  webkit_settings_set_javascript_can_access_clipboard(settings, true);
  webkit_settings_set_enable_html5_local_storage(settings, true);
  webkit_settings_set_hardware_acceleration_policy(settings, WEBKIT_HARDWARE_ACCELERATION_POLICY_ALWAYS);
  webkit_settings_set_enable_write_console_messages_to_stdout(settings, true);
  webkit_settings_set_enable_developer_extras(settings, true);

  // Load target URL
  webkit_web_view_load_uri(WEBKIT_WEB_VIEW(webview), APP_URL);

  // FullScreen or Maximize window
  if (APP_DISPLAY_MODE == "fullscreen"){
    gtk_window_fullscreen(GTK_WINDOW(window));
  } else if (APP_DISPLAY_MODE == "maximized"){
    gtk_window_maximize(GTK_WINDOW(window));
  }

  // Show window
  gtk_widget_grab_focus(GTK_WIDGET(webview));
  gtk_window_present(GTK_WINDOW(window));
}

// On shutdown callback
static void on_shutdown_app(GtkApplication *app)
{
}

// Start application
int start_application(int argc, char *argv[])
{

  // Create a new application
  GtkApplication *app = gtk_application_new(APP_ID, G_APPLICATION_DEFAULT_FLAGS);

  // Attach signal callbacks
  g_signal_connect(app, "startup", G_CALLBACK(on_startup_app), NULL);
  g_signal_connect(app, "activate", G_CALLBACK(on_activate_app), NULL);
  g_signal_connect(app, "shutdown", G_CALLBACK(on_shutdown_app), NULL);

  // Run application
  int status = g_application_run(G_APPLICATION(app), argc, argv);
  g_object_unref(app);

  return status;
}