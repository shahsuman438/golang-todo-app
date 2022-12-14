package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Completed bool               `json:"completed,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
}
