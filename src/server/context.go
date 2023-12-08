package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context struct
type Context struct {
	URI        string
	Request    *http.Request
	Response   http.ResponseWriter
	StatusCode int
}

// Bind body JSON data into destination JSON struct
func (c *Context) Bind(destination any) error {
	err := json.NewDecoder(c.Request.Body).Decode(&destination)
	if err != nil {
		return err
	}

	return nil
}

// Header set the HTTP response header
func (c *Context) Header(key string, value string) {
	c.Response.Header().Set(key, value)
}

// Status set the status code for HTTP response
func (c *Context) Status(code int) *Context {
	c.StatusCode = code
	return c
}

// Error output response error
func (c *Context) Error(err error) error {
	http.Error(c.Response, err.Error(), c.StatusCode)
	return nil
}

// String output response content as string
func (c *Context) String(content string) error {
	c.Header("Content-Type", "text/plain; charset=utf-8")
	c.Response.WriteHeader(c.StatusCode)
	fmt.Fprintln(c.Response, content)
	return nil
}

// JSON output response content as JSON
func (c *Context) JSON(content any) error {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Response.WriteHeader(c.StatusCode)
	return json.NewEncoder(c.Response).Encode(content)
}
