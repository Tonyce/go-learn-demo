package model

import (
	"context"
	"fmt"
	"logical-example/internal/common"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User ...
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Password string             `bson:"password" json:"password"`
	Roles    []string           `bson:"roles" json:"roles"`
}

// UserRepository 用户仓库
type UserRepository interface {
	// 根据用户id查询得到一个用户或是错误信息
	FindOne(id int) (*User, error)
}

type mongoUserRepository struct {
	collection *mongo.Collection
}

// MongoUserRepository ...
var MongoUserRepository UserRepository

func init() {
	fmt.Println("model user init")
	mongoClient := common.GetMongoClient()
	db := mongoClient.Database("test")
	collection := db.Collection("user")
	MongoUserRepository = &mongoUserRepository{
		collection: collection,
	}
}

// FindOne ...
func (repo *mongoUserRepository) FindOne(id int) (*User, error) {
	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := repo.collection.FindOne(ctx, bson.D{})
	err := result.Decode(&user)

	return &user, err
}
