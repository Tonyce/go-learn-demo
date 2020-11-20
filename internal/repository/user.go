package repository

import (
	"context"
	"fmt"
	"logical-example/internal/common"
	"logical-example/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// UserRepository 用户仓库
type UserRepository interface {
	// 根据用户id查询得到一个用户或是错误信息
	// FindOne(id int) (*model.User, error)
	// FindOneFromRedis(id int) (*model.User, error)
	mongoUserRepository
	redisUserRepository
}

type mongoUserRepository interface {
	FindOne(id int) (*model.User, error)
}

type redisUserRepository interface {
	FindOneFromRedis(id int) (*model.User, error)
}

type user struct {
	// mongo      mongoUserRepository
	// redis      redisUserRepository
	collection *mongo.Collection
}

// UserRepo ...
var UserRepo UserRepository

func init() {
	fmt.Println("user reposity init")
	mongoClient := common.DB.MongoClient
	db := mongoClient.Database("test")
	collection := db.Collection("user")
	UserRepo = &user{
		collection: collection,
	}
}

// FindOne ...
func (repo *user) FindOne(id int) (*model.User, error) {
	var user model.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := repo.collection.FindOne(ctx, bson.M{})
	err := result.Decode(&user)

	return &user, err
}

// FindOneFromRedis
func (repo *user) FindOneFromRedis(id int) (*model.User, error) {
	return &model.User{}, nil
}
