package main

import "net/http"


func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write a simple JSON response indicating the service is healthy
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"status": "healthy"}`))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("OK..."))
}