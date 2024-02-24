package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"todo_mongo/service"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

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
    var todo service.Todo

    w.Header().Set("Content-Type", "application/json")

    todos, err := todo.GetAllTodos()
    if err != nil {
        log.Println(err)
        return 
    }

    json.NewEncoder(w).Encode(todos)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
    var todo service.Todo

    id := chi.URLParam(r, "id")
    fmt.Println("ID -> ", id)

    w.Header().Set("Content-Type", "application/json")

    todo, err := todo.GetTodoById(id)
    if err != nil {
        log.Println(err)
        return
    }

    fmt.Println(todo)

    json.NewEncoder(w).Encode(todo)
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
