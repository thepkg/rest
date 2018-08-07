package request

import (
	"net/http"
)

// Handler describes default handler function.
type Handler func(http.ResponseWriter, *http.Request)
