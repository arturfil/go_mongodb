package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"todo_mongo/db"
	"todo_mongo/handler"
	"todo_mongo/service"

)

type Application struct {
    Models service.Todo
}

func main() {
    mongoClient, err := db.ConnectToMongo()
    if err != nil {
        log.Panic(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
    defer cancel()

    defer func() {
        if err = mongoClient.Disconnect(ctx); err != nil {
            panic(err)
        }
    }()

    service.New(mongoClient)

    log.Println("Server running in port", 8080)
	log.Fatal(http.ListenAndServe(":8080", handler.CreateRouter()))
    
}

