-- Disconnect from any open database
\c postgres

-- Drop the database if it already exists
DROP DATABASE IF EXISTS youtube_db;

-- Create the database
CREATE DATABASE youtube_db;

-- Connect to the database
\c youtube_db;

-- Create the videos table if not exists
CREATE TABLE IF NOT EXISTS videos (
                                      id SERIAL PRIMARY KEY,
                                      title TEXT,
                                      video_url TEXT,
                                      upload_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Grant all privileges on the table
GRANT ALL PRIVILEGES ON TABLE videos TO users;
