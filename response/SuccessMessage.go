package response

const (
	// SuccessMessageFormat describes how must looks success message JSON payload.
	SuccessMessageFormat = `{"error":null,"success":{"code":%d,"data":%s}}`
)

// SuccessMessage describes success HTTP response message.
type SuccessMessage struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
