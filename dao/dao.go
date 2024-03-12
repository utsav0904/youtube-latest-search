package dao

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitPostgreSQL(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	log.Println("Connected to PostgreSQL database")
	return nil
}

func InsertVideo(title, videoURL, uploadDate string) error {
	// Check if the videos table exists
	rows, err := db.Query("SELECT to_regclass('public.videos')")
	if err != nil {
		return err
	}
	defer rows.Close()

	var tableName sql.NullString
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			return err
		}
	}

	// If tableName is null, the videos table does not exist; create it
	if !tableName.Valid || tableName.String == "" {
		if err := createVideosTable(); err != nil {
			return err
		}
	}

	// Insert the video data
	_, err = db.Exec("INSERT INTO videos (title, video_url, upload_date) VALUES ($1, $2, $3)", title, videoURL, uploadDate)
	if err != nil {
		return err
	}
	return nil
}

func createVideosTable() error {
	_, err := db.Exec(`CREATE TABLE videos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		video_url TEXT,
		upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return err
	}
	log.Println("Created videos table")
	return nil
}

func GetSortedVideos(limit int) ([]map[string]string, error) {
	rows, err := db.Query("SELECT title, video_url, upload_date FROM videos ORDER BY upload_date DESC LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []map[string]string
	for rows.Next() {
		var title, videoURL, uploadDate string
		if err := rows.Scan(&title, &videoURL, &uploadDate); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		video := map[string]string{
			"title":       title,
			"video_url":   videoURL,
			"upload_date": uploadDate,
		}
		videos = append(videos, video)
	}

	return videos, nil
}
