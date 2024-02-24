BINARY=mon_Go

build:
	go build -o ${BINARY} cmd/main.go 

run:
	./${BINARY}

restart: build run
