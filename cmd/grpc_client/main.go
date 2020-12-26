/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"logical-example/internal/pb"
	"os"
	"time"

	// pb "logical-example/internal/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

const (
	address     = "localhost:5055"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	go func() {
		count := 0
		state := conn.GetState()
		for {
			v := conn.WaitForStateChange(context.Background(), state)
			state = conn.GetState()
			if count > 5 {
				conn.Close()
				fmt.Println("close")
				break
			}
			fmt.Println("--------------", v)

			switch state {
			case connectivity.Connecting:
				{
					fmt.Println("connneting")
					count++
				}
			case connectivity.Ready:
				{

					fmt.Println("Ready")
					count = 0
				}
			case connectivity.Shutdown:
				{
					fmt.Println("Shutdown")
				}
			case connectivity.TransientFailure:
				{
					fmt.Println("TransientFailure")
				}
			}
		}
	}()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		time.Sleep(60 * time.Minute)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	time.Sleep(60 * time.Minute)
}
