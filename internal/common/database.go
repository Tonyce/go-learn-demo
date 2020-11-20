package common

import (
	"fmt"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

// Database ...
type Database struct {
	MongoClient *mongo.Client
	RedisClient *redis.Client
}

// DB ...
var DB *Database

func init() {
	fmt.Println("db init")
	DB = initDB()
}

func initDB() *Database {
	initMongo()
	initRedis()
	return &Database{
		MongoClient: mdb,
		RedisClient: rdb,
	}
}
