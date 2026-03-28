package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `json:"name" validate:"required,min=2"`
	Email    string             `json:"email" validate:"email,required"`
	Password string             `json:"password" validate:"required,min=6"`
	Role     string             `json:"role" validate:"required,oneof=user admin"`
}

type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `json:"name" validate:"required"`
	Price       float64            `json:"price" validate:"required"`
	Description string             `json:"description"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}