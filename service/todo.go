package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

// Struct where we are going to be inserting objects with information
type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool      `json:"completed,omitempty" bson:"completed,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

func (t *Todo) InsertTodo(entry Todo) error {
    collection := client.Database("todosdb").Collection("todos")

    _, err := collection.InsertOne(context.TODO(), Todo{
        Task: entry.Task,
        Completed: entry.Completed,
        CreatedAt: time.Now(),
    })
    if err != nil {
        log.Println("Error inserting obj", err) 
        return err
    }

    return nil
}
