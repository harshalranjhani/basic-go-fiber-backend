package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty"`
	Name string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Title string `json:"title" validate:"required"`
}
