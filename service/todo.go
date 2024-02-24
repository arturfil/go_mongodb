package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func New(mongo *mongo.Client) Todo {
	client = mongo

	return Todo{}
}

// Struct where we are going to be inserting objects with information
type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool      `json:"completed,omitempty" bson:"completed,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"createdAt,omitempty"`
	UpatedAt  time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func returnCollectionPointer() *mongo.Collection {
    return client.Database("todosdb").Collection("todos")
}

func (t *Todo) GetAllTodos() ([]Todo, error) {
	collection := returnCollectionPointer()
	var todos []Todo

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *Todo) GetTodoById(id string) (Todo, error){
    collection := returnCollectionPointer()
    var todo Todo

    mongoID, err := primitive.ObjectIDFromHex(id)

    err = collection.FindOne(context.Background(), bson.M{"_id": mongoID}).Decode(&todo)
    if err != nil {
        return Todo{}, err
    }

    return todo, nil

}

func (t *Todo) InsertTodo(entry Todo) error {
	collection := client.Database("todosdb").Collection("todos")

	_, err := collection.InsertOne(context.TODO(), Todo{
		Task:      entry.Task,
		Completed: entry.Completed,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting obj", err)
		return err
	}

	return nil
}

func (t *Todo) UpdateTodo() {

}

func (t *Todo) DeleteTodo() {

}
