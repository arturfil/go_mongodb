package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Todo structure
type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"task,omitempty"`
	Completed bool      `json:"completed,omitempty" bson:"completed,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}

var collection *mongo.Collection

// Connect to MongoDB
func connectToMongo() {
	// MongoDB connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the collection
	collection = client.Database("tododb").Collection("todos")
}

// Get all todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todos []Todo
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	json.NewEncoder(w).Encode(todos)
}

// Get a single todo
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	var todo Todo
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(todo)
}

// Create a new todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todo.CreatedAt = time.Now()

	result, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

// Update a todo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	result, err := collection.ReplaceOne(context.Background(), bson.M{"_id": id}, todo)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

// Delete a todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	// Connect to MongoDB
	connectToMongo()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Endpoints
	r.Get("/todos", getTodos)
	r.Get("/todos/{id}", getTodo)
	r.Post("/todos", createTodo)
	r.Put("/todos/{id}", updateTodo)
	r.Delete("/todos/{id}", deleteTodo)

	// Server
	log.Fatal(http.ListenAndServe(":8080", r))
}

