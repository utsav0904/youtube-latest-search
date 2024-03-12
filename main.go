package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test-youtube/dao"
	"test-youtube/handler"
)

const (
	PostgreSQLDataSource = "host=db port=5432 user=users password=password dbname=youtube_db sslmode=disable"
	ServerPort           = ":8080"
)

func main() {
	if err := dao.InitPostgreSQL(PostgreSQLDataSource); err != nil {
		log.Fatal("Error initializing PostgreSQL:", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/search_videos", handler.SearchVideos).Methods("GET")
	r.HandleFunc("/get_sorted_videos", handler.GetSortedVideosHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
