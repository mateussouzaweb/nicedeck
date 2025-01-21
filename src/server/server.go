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
		uri = filepath.ToSlash(uri)

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

		// Grab 404 route when handle not found
		// Must ensure that 404 will always be found
		if handler == nil {
			handler = Match("GET", "/404")
		}

		// Run handler
		err := handler(context)
		if err != nil {
			context.Status(http.StatusInternalServerError).Error(err)
		}
	})

	// Initiate listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	// Attach server to listener
	server := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  0 * time.Second,
		WriteTimeout: 0 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	ready <- true
	return server.Serve(listener)
}
