package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// SUPPORTED FILES
	tif = ".tif"
	csv = ".csv"

	// HEADERS
	contentType = "Content-Type"
	appJSON     = "application/json; charset=UTF-8"

	// ERRORS
	connectionError      = "Failed to connect %v"
	errorReadingJSON     = "Error reading JSON data."
	errorReadingFilename = "Unsuported filename."
	errorReadingFileData = "Error reading data image %v"
	errorRetrievingData  = "Error retrieving %v image."

	// SUCCESS MESSAGES
	bWritten = "Bytes written: %d. "
	bRead    = "Bytes read: %d. "
)

// jsonResponse builds the endpoints response
func jsonResponse(w http.ResponseWriter, statusCode int, response ...interface{}) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
