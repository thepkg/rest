package rest

import (
	"encoding/json"
	"fmt"
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

// MustUnmarshalBody performs read and json.Unmarshal body into data.
func MustUnmarshalBody(r *http.Request, data interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(fmt.Errorf("failed to read request body, error: %s", err))
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal request body, error: %s", err))
	}
}
