package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Name     string             `bson:"name"`
	Age      int8               `bson:"age"`
}
