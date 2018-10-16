package rest

import (
	"net/http"

	"github.com/thepkg/rest/request"
)

var (
	handlerRegistry = request.HandlerRegistry{}
)

func init() {
	handlerRegistry = request.NewHandlerRegistry()
}

// HEAD registers the HEAD HTTP method handler function for the given pattern.
func HEAD(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handlerRegistry.Handle("HEAD", pattern, handler)
}

// GET registers the GET HTTP method handler function for the given pattern.
func GET(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handlerRegistry.Handle("GET", pattern, handler)
}

// POST registers the POST HTTP method handler function for the given pattern.
func POST(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handlerRegistry.Handle("POST", pattern, handler)
}

// PATCH registers the PATCH HTTP method handler function for the given pattern.
func PATCH(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handlerRegistry.Handle("PATCH", pattern, handler)
}

// PUT registers the PUT HTTP method handler function for the given pattern.
func PUT(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handlerRegistry.Handle("PUT", pattern, handler)
}

// DELETE registers the DELETE HTTP method handler function for the given pattern.
func DELETE(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handlerRegistry.Handle("DELETE", pattern, handler)
}

// OPTIONS registers the OPTIONS HTTP method handler function for the given pattern.
func OPTIONS(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	handlerRegistry.Handle("OPTIONS", pattern, handler)
}
