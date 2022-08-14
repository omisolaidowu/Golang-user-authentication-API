package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	First_name string             `bson:"First_name,omitempty bson:"First_name,omitempty"`
	Last_name  string             `bson:"Last_name,omitempty bson:"Last_name,omitempty"`
	Username   string             `bson:"Username,omitempty bson:"Username,omitempty"`
	Email      string             `bson:"Email,omitempty bson:"Email,omitempty"`
	Password   string             `bson:"Password,omitempty bson:"Password,omitempty"`
}
