package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"test-youtube/controller"
	"time"
)

// Define a global channel to signal the background task to stop
var stopBackgroundTask = make(chan struct{})
var isBackgroundTaskRunning bool

func GetSortedVideosHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Search query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	// Check if background task is already running
	if !isBackgroundTaskRunning {
		// Start background task to fetch and store videos
		go startBackgroundTask(query)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Background task started or is already running to fetch and store videos."))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Background task started to fetch and store videos."))

	limit := 10 // Change the limit as per your requirement

	videos, err := controller.GetSortedVideoController(limit, query)
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

// Start background task to continuously fetch and store videos
func startBackgroundTask(query string) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call controller function to fetch latest videos and store data
			if err, _ := controller.SearchController(query); err != nil {
				log.Println("Failed to fetch and store videos:", err)
			}
		case <-stopBackgroundTask:
			return
		}
	}
}
