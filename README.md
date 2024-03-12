Steps to build and test

1. Clone the repo locally
2. Go to repo directory
3. Compose and build docker using below command\
```docker compose up --build -d```
4. Once done, open postman and use below curl to hit API\
    ```curl -X GET 'http://localhost:8080/search_videos?q=football'```
5. Use below command to check database\
```docker exec -it my-postgres psql -U users youtube_db```

