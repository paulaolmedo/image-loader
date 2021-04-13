package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// jsonResponse .
func jsonResponse(w http.ResponseWriter, statusCode int, response interface{}) {

	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(response) //puede ser datos de verdad o un error

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error()) //error parseando el mensaje
	}

}
