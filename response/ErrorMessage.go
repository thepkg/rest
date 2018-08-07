package response

const (
	// ErrorMessageFormat describes how must looks error message JSON payload.
	ErrorMessageFormat = `{"error":{"code":%d,"data":%s},"success":null}`
)

// ErrorMessage describes error HTTP response message.
type ErrorMessage struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
