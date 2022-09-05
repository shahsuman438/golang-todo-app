package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"todoapp/configs"
	"todoapp/models"
	"todoapp/services"

	"github.com/go-playground/validator/v10"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection = configs.GetCollection(configs.DB, "todo")
var validate = validator.New()

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var todo models.Todo
		//validate the request body
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&todo); validationErr != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		result := services.CreateTodo(&todo)
		jsonResp, err := json.Marshal(result)
		w.Write(jsonResp)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		fmt.Println("id:-", id)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		result, err := services.GetATodo(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		jsonResp, err := json.Marshal(result)
		w.Write(jsonResp)
		fmt.Println("data", result)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func EditATodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id := r.URL.Query().Get("id")
		fmt.Println()
		var todo models.Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		update := bson.M{"title": todo.Title, "completed": todo.Completed}
		result, err := services.EditATodo(id, update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(result)
		w.Write(jsonResp)
		fmt.Println("data", jsonResp)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func DeleteATodo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		fmt.Println("id:-", id)
		w.Header().Set("Content-Type", "application/json")
		result, err := services.DeleteATodo(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if result.DeletedCount < 1 {
			http.Error(w, "Todo Not Found", http.StatusNotFound)
			return
		}
		jsonResp, err := json.Marshal(result)
		w.Write(jsonResp)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		results, err := services.GetAllTodo()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(results)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(jsonResp)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
