build:
	docker build -t doge-docker .

make run:
	docker run -p 3000:3000 --name=go_doge doge-docker

getRecords:
	curl http://localhost:3000/v1/getRecords

make stop:
	docker rm -f go_doge