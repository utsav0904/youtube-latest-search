package controller

import (
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"test-youtube/dao"
)

const APIKey = "AIzaSyD_a-Epucnj_PwOXo3lK6PZ1mDybPhquAo"

func SearchController(query string) ([]map[string]string, error) {
	service, err := youtube.New(&http.Client{
		Transport: &transport.APIKey{Key: APIKey},
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
		}

		if err := dao.InsertVideo(video["title"], video["video_url"], video["upload_date"]); err != nil {
			log.Println("Error inserting video:", err)
		}

		videos = append(videos, video)
	}

	return videos, nil
}

func GetSortedVideoController(limit int) ([]map[string]string, error) {
	videos, err := dao.GetSortedVideos(limit)
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
