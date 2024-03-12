package handler

import (
	"encoding/json"
	"net/http"
	"test-youtube/controller"
)

func SearchVideos(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Search query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	// Call controller function to handle search logic
	videos, err := controller.SearchController(query)
	if err != nil {
		http.Error(w, "Failed to perform YouTube search", http.StatusInternalServerError)
		return
	}

	// Convert videos to JSON response
	jsonResponse, err := json.Marshal(videos)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetSortedVideosHandler(w http.ResponseWriter, r *http.Request) {
	limit := 10 // Change the limit as per your requirement

	videos, err := controller.GetSortedVideoController(limit)
	if err != nil {
		http.Error(w, "Failed to retrieve sorted videos", http.StatusInternalServerError)
		return
	}

	// Convert videos to JSON response
	jsonResponse, err := json.Marshal(videos)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
