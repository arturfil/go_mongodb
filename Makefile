include .env

BINARY=mon_Go

build:
	go build -o ${BINARY} cmd/main.go 

up:
	docker-compose up --build -d --remove-orphans

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

post.get_all:
	@echo "Getting all posts..."
	curl -X GET http://localhost:8080/api/todos | json_pp

post.create:
	curl -X POST http://localhost:8080/api/todos/create \
	-H "Content-Type: application/json" \
	-d \
	'{ \
		"task": "Do commentary video leetcode video", \
		"completed": false \
	}' \



run:
	@env MONGO_DB_USERNAME=${MONGO_DB_USERNAME} MONGO_DB_PASSWORD=${MONGO_DB_PASSWORD} ./${BINARY}

restart: build run
