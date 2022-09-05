package services

import (
	"context"
	"time"
	"todoapp/configs"
	"todoapp/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection = configs.GetCollection(configs.DB, "todo")

func CreateTodo(todo *models.Todo) any {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	newTodo := models.Todo{
		Id:        primitive.NewObjectID(),
		Completed: todo.Completed,
		Title:     todo.Title,
	}
	result, err := todoCollection.InsertOne(ctx, newTodo)
	if err != nil {
		return err
	}
	return result
}

func GetATodo(todoId string) (*models.Todo, error) {
	var todo models.Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(todoId)
	err := todoCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func EditATodo(id string, todo primitive.M) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(id)
	_, err := todoCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": todo})
	if err != nil {
		return nil, err
	}
	
	result, err := GetATodo(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DeleteATodo(todoId string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(todoId)
	result, err := todoCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllTodo() ([]models.Todo, error) {
	var todos []models.Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	results, err := todoCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleTodo models.Todo
		if err = results.Decode(&singleTodo); err != nil {
			return nil, err
		}

		todos = append(todos, singleTodo)
	}
	return todos, nil
}
