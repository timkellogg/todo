package controllers

import (
	"net/http"
)

// SetHeaders sets common headers needed for routes such as content-type
func SetHeaders(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json charset=UTF-8")
	return w
}

// HandleError : handles the error by logging
func HandleError(err error) {
	panic(err)
}
