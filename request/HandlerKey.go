package request

// HandlerKey describes compound handler key which is used in HandlerRegistry.
type HandlerKey string

// MakeKey makes HandlerKey.
func MakeKey(method string, pattern string) HandlerKey {
	return HandlerKey(method + ":" + pattern)
}
