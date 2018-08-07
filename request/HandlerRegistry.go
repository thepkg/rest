package request

import (
	"net/http"

	"../response"
)

// HandlerRegistry this is simple proxy between top level application and native "net/http" package.
// This registry holds map between HTTP method + pattern (URL path) and handler function,
// and performs routing for HTTP requests.
type HandlerRegistry struct {
	handlers      map[HandlerKey]Handler
	boundPatterns map[string]bool
}

// NewHandlerRegistry returns new instance of HandlerRegistry.
func NewHandlerRegistry() HandlerRegistry {
	hr := HandlerRegistry{}
	hr.handlers = make(map[HandlerKey]Handler)
	hr.boundPatterns = make(map[string]bool)

	return hr
}

// Handle performs next steps:
// 1) add handler into HandlerRegistry.
// 2) bind pattern (URL path) to HandlerRegistry.
func (hr HandlerRegistry) Handle(method string, pattern string, handler Handler) {
	key := MakeKey(method, pattern)
	hr.add(key, handler)
	hr.bind(pattern)
}

func (hr HandlerRegistry) add(key HandlerKey, handler Handler) {
	hr.handlers[key] = handler
}

func (hr HandlerRegistry) isBound(pattern string) bool {
	enabled, exists := hr.boundPatterns[pattern]

	return exists && enabled
}

func (hr HandlerRegistry) bind(pattern string) {
	if hr.isBound(pattern) {
		return
	}

	hr.boundPatterns[pattern] = true

	http.HandleFunc(pattern, func(res http.ResponseWriter, req *http.Request) {
		key := MakeKey(req.Method, pattern)
		handler, exists := hr.handlers[key]
		if exists {
			handler(res, req)
		} else {
			message := response.Message{
				Error:   &response.ErrorMessage{Code: 405, Data: "Method Not Allowed"},
				Success: nil,
			}
			response.JSON(res, 405, message)
		}
	})
}
