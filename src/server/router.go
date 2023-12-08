package server

import (
	"regexp"
	"strings"
)

// Handle function
type Handler func(context *Context) error

// Middleware function
type Middleware = func(next Handler) Handler

// Route struct
type Route struct {
	Method   string
	Endpoint string
	Handler  Handler
}

var routes []*Route
var middlewares []Middleware

// Use middleware in routers
func Use(middleware Middleware) {
	middlewares = append(middlewares, middleware)
}

// Add route to server
func Add(method string, endpoint string, handler Handler) {

	// Apply middlewares
	for i := range middlewares {
		handler = middlewares[len(middlewares)-1-i](handler)
	}

	// Append route
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

		// Single endpoint matching
		if endpoint == route.Endpoint {
			return route.Handler
		}

		// Regex endpoint matching
		endpointPattern := strings.ReplaceAll(route.Endpoint, "/", "\\/")
		endpointRegex := regexp.MustCompile(endpointPattern)

		if endpointRegex.MatchString(endpoint) {
			return route.Handler
		}

	}

	return nil
}
