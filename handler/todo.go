package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo_mongo/service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Msg:  "Health Check",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)
}

// GetTodos - Returns all the todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("todosdb").Collection("todos")

	var todos []service.Todo
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
	    log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
	    var todo service.Todo
	    cursor.Decode(&todo)
	    todos = append(todos, todo)
	}

    fmt.Println("TODOS", todos)

	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")

    var todo service.Todo

    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        log.Fatal("error", err) 
    }

    err = todo.InsertTodo(todo)
    if err != nil {
        errorRes := Response{
            Msg: "Error",
            Code: 304,
        }
        json.NewEncoder(w).Encode(errorRes)
        return 
    }
   
	res := Response{
		Msg:  "Success",
		Code: 201,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)

}
