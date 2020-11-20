package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User ...
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Password string             `bson:"password" json:"password"`
	Roles    []string           `bson:"roles" json:"roles"`
}
