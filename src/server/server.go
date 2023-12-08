package server

import (
	"net"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

// Start the server with given address.
// Server will serve static resource files by default.
// Add routes before starting it for more endpoints
func Start(address string, ready chan bool) error {

	// Attach server handle
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {

		// Get clean URI
		uri := filepath.Clean(request.URL.Path)
		if len(uri) > 1 {
			uri = strings.TrimSuffix(uri, "/")
		}

		// Create context
		context := &Context{
			URI:        uri,
			Request:    request,
			Response:   response,
			StatusCode: http.StatusOK,
		}

		// Check for matching endpoints
		handler := Match(request.Method, uri)

		// When found match, run handler
		if handler != nil {
			err := handler(context)
			if err != nil {
				context.Status(http.StatusInternalServerError).Error(err)
			}
			return
		}

		// Return 404 if not found
		context.Status(http.StatusNotFound).String(http.StatusText(http.StatusNotFound))
	})

	// Initiate listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// Listener is ready
	ready <- true

	// Attach server to listener
	server := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server.Serve(listener)
}
