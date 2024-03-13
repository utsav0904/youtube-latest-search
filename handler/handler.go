package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"test-youtube/controller"
	"time"
)

// Define a global channel to signal the background task to stop
var stopBackgroundTask = make(chan struct{})
var isBackgroundTaskRunning bool

func GetSortedVideosHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for pagination
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page <= 0 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	limit := 10 // Change the limit as per your requirement
	offset := (page - 1) * limit

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

	// Fetch paginated videos
	videos, err := controller.GetSortedVideoController(limit, offset, query)
	if err != nil {
		http.Error(w, "Failed to retrieve paginated sorted videos", http.StatusInternalServerError)
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
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Call controller function to fetch latest videos and store data
			if _, err := controller.SearchController(query); err != nil {
				log.Println("Failed to fetch and store videos:", err)
			}
		case <-stopBackgroundTask:
			return
		}
	}
}
