package main

import (
	"fmt"
	"logical-example/internal/apiserver"
)

func init() {
	fmt.Println("main init")
}

func main() {
	server := apiserver.SetupRouter() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	server.Run()
}
