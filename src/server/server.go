package server

import (
	"errors"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/mateussouzaweb/nicedeck/frontend"
)

// Start the server with given address.
// Server will serve static resource files by default.
// Add routes before starting it for more endpoints
func Start(address string) error {

	// Open fs and http handle for static content
	staticFS := frontend.GetStaticFS()
	staticHandler := http.FileServer(http.FS(staticFS))

	// Attach server handle
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		uri := filepath.Clean(request.URL.Path)
		uri = strings.TrimSuffix(uri, "/")

		// Create context
		context := &Context{
			Request:    request,
			Response:   response,
			StatusCode: http.StatusOK,
		}

		// Check for custom endpoints
		// Implementation is simple, but works for the propurse on this application
		handler := Match(request.Method, uri)

		// When match, run handler
		if handler != nil {
			err := handler(context)
			if err != nil {
				context.Status(http.StatusInternalServerError).Error(err)
			}
			return
		}

		// Endpoint was not found, so we try to find from static files
		// First, make sure we are requesting a file when trying to get an unknown uri
		if filepath.Ext(uri) == "" {
			uri += "/index.html"
		}

		// Check if file exist as directories where discarded in previous line
		// NOTE: uri will always make sure it is a file, so we do not need to check for
		file, err := staticFS.Open(strings.TrimPrefix(uri, "/"))
		if err != nil {
			// Not found when file is not detected
			if errors.Is(err, fs.ErrNotExist) {
				context.Status(http.StatusNotFound).String(http.StatusText(http.StatusNotFound))
				return
			}

			// Server error when are other type of error
			context.Status(http.StatusInternalServerError).Error(err)
			return
		}

		// Make sure file is closed
		defer file.Close()

		// Server static files
		staticHandler.ServeHTTP(context.Response, context.Request)
	})

	return http.ListenAndServe(address, mux)
}
