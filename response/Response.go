package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON sends HTTP response with JSON payload.
func JSON(w http.ResponseWriter, code int, message Message) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(message)
	if err != nil {
		w.WriteHeader(500)
		er := fmt.Errorf("failed to perform JSON encode, error: %s", err)
		fmt.Fprintf(w, ErrorMessageFormat, 500, er)
	}
}
