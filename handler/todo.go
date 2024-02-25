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
var todo service.Todo

// HealthCheck - returns json message to verify the api works
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
    todos, err := todo.GetAllTodos()
    if err != nil {
        log.Println(err)
        return 
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

// GetTodoById - returns the todo, provided an id 
func GetTodoById(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    todo, err := todo.GetTodoById(id)
    if err != nil {
        log.Println(err)
        return
    }

    fmt.Println(todo)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

// CreateTodo - Inserts a Todo obect in the db
func CreateTodo(w http.ResponseWriter, r *http.Request) {
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
		Code: http.StatusOK,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonStr)

}

// UpdateTodo - udpates a preexisting todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
    err := json.NewDecoder(r.Body).Decode(&todo)
    if err != nil {
        log.Println(err)
        return 
    }
    
    _, err = todo.UpdateTodo(todo)

    res := Response{
        Msg: "Successfully Updated",
        Code: 202,
    }

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

// DeleteTodo - deletes the todo from the db
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")

    err := todo.DeleteTodo(id)
    if err != nil {
        log.Println(err)
        return 
    }

    res := Response{
        Msg: "Successfully Deleted",
        Code: 204,
    }

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(res.Code)
    // redundant, 204: successful deletion with no content
	w.Write(jsonStr)

}
