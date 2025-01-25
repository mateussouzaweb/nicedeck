package server

import (
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/mateussouzaweb/nicedeck/src/cli"
)

// Start the server with given address and listener.
// Server will serve static resource files by default.
// Add routes before starting it for more endpoints
func Start(address string, listener net.Listener, ready chan bool) error {

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

	// Attach server to listener
	server := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  0 * time.Second,
		WriteTimeout: 0 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		ready <- true
	}()

	return server.Serve(listener)
}

// Init server
func Init(version string, developmentMode bool, address string, ready chan bool, shutdown chan bool) error {

	// Initiate listener when possible
	addressInUse := false
	listener, err := net.Listen("tcp", address)

	if err != nil {
		if err, ok := err.(*net.OpError); ok {
			if err, ok := err.Err.(*os.SyscallError); ok {
				if err.Err == syscall.EADDRINUSE {
					addressInUse = true
				}
			}
		}
		if !addressInUse {
			return err
		}
	}

	// If address already is in use, skip server init
	if addressInUse {
		cli.Printf(cli.ColorWarn, "Server already is running from another instance\n")
		cli.Printf(cli.ColorWarn, "Skipping server startup and closing current process...\n")
		ready <- true

		time.Sleep(1 * time.Second)
		shutdown <- true
		return nil
	} else {
		defer listener.Close()
	}

	// Setup the server with shutdown channel
	err = Setup(version, developmentMode, shutdown)
	if err != nil {
		return err
	}

	// Start the server and wait for requests
	err = Start(address, listener, ready)
	if err != nil {
		return err
	}

	return nil
}
