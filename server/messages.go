package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// jsonResponse builds the endpoints response
func jsonResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
