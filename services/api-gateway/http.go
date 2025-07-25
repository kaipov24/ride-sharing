package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if reqBody.UserID == "" {
		http.Error(w, "UserID is required", http.StatusBadRequest)
		return
	}

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)

	resp, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	if err != nil {
		http.Error(w, "Failed to call trip service", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var responseBody any
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		http.Error(w, "Failed to decode response from trip service", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: responseBody}

	writeJSON(w, http.StatusCreated, response)
}
