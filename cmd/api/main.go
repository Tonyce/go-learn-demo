package main

import (
	"fmt"
	"log"
	"logical-example/internal/apiserver"
)

func init() {
	fmt.Println("main init")
}

func main() {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// err = client.Connect(ctx)
	// common.InitMongoDB()

	server := apiserver.SetupRouter() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
