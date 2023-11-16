package server

import (
	"strings"
)

// Handle function
type Handler func(context *Context) error

// Route struct
type Route struct {
	Method   string
	Endpoint string
	Handler  Handler
}

var routes []*Route

// Add route to server
func Add(method string, endpoint string, handler Handler) {
	routes = append(routes, &Route{
		Method:   strings.ToUpper(method),
		Endpoint: endpoint,
		Handler:  handler,
	})
}

// Match route with given method and endpoint and return its handler
func Match(method string, endpoint string) Handler {

	// Make sure HEAD is GET
	if method == "HEAD" {
		method = "GET"
	}

	// Find matching route
	for _, route := range routes {

		// Check HTTP method
		if method != route.Method {
			continue
		}

		// Check endpoint matching
		if endpoint != route.Endpoint {
			continue
		}

		// When match, run handler
		return route.Handler
	}

	return nil
}
