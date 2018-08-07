package response

// Message describes default HTTP message (both types: error and success).
type Message struct {
	Error   *ErrorMessage   `json:"error"`
	Success *SuccessMessage `json:"success"`
}
