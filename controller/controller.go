package controller

import (
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
	"test-youtube/dao"
)

func SearchController(query string) ([]map[string]string, error) {
	log.Println("Fetch Videos API Running in background every 20 Seconds")
	service, err := youtube.New(&http.Client{
		Transport: &transport.APIKey{Key: os.Getenv("API_KEY")},
	})
	if err != nil {
		return nil, err
	}

	call := service.Search.List([]string{"id", "snippet"}).
		Q(query).
		Type("video").
		MaxResults(10)

	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	var videos []map[string]string
	for _, item := range response.Items {
		video := map[string]string{
			"title":         item.Snippet.Title,
			"upload_date":   item.Snippet.PublishedAt,
			"thumbnail_url": item.Snippet.Thumbnails.Default.Url,
			"video_url":     "https://www.youtube.com/watch?v=" + item.Id.VideoId,
			"query":         query,
		}

		if err := dao.InsertVideo(video["title"], video["video_url"], video["upload_date"], video["query"]); err != nil {
			log.Println("Error inserting video:", err)
		}

		videos = append(videos, video)
	}
	log.Println("New Videos added to videos table")

	return videos, nil
}

func GetSortedVideoController(limit int, query string) ([]map[string]string, error) {
	videos, err := dao.GetSortedVideos(limit, query)
	if err != nil {
		log.Println("Error retrieving sorted videos:", err)
		return nil, err
	}

	var sortedVideos []map[string]string
	for _, video := range videos {
		sortedVideo := map[string]string{
			"title":       video["title"],
			"video_url":   video["video_url"],
			"upload_date": video["upload_date"],
		}
		sortedVideos = append(sortedVideos, sortedVideo)
	}

	return sortedVideos, nil
}
