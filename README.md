Steps to build and test

1. Clone the repo locally
2. Go to repo directory
3. Compose and build docker using below command\
```docker compose up --build -d```
4. Once done, open postman and use below curl to hit API\
    ```curl -X GET 'http://localhost:8080/get_sorted_videos?q=football&page=1'```
5. Use below command to check database\
```docker-compose exec db psql -U users youtube_db```
6. Check logs for app and db by below commands
```docker-compose logs app/db```
