build:
	docker-compose up -d --rm

getRecords:
	docker exec -it godoge_doge_1 curl http://localhost:3000/v1/getRecords