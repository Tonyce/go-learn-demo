package controller

import (
	"context"
	"log"
	"logical-example/internal/apiserver/service"
	"logical-example/internal/pb"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

var conn *grpc.ClientConn

func init() {
	r := manual.NewBuilderWithScheme("whatever")
	r.InitialState(resolver.State{Addresses: []resolver.Address{
		// {Addr: "127.0.0.1:5050"},
		{Addr: "127.0.0.1:5051"},
		{Addr: "127.0.0.1:5052"},
	}})
	// cc, err := grpc.Dial(
	// 	r.Scheme()+":///test.server",
	// 	grpc.WithInsecure(),
	// 	grpc.WithResolvers(r),
	// 	grpc.WithBalancerName(roundrobin.Name)
	// )
	var err error
	conn, err = grpc.Dial(
		r.Scheme()+":///test.server",
		grpc.WithInsecure(),
		grpc.WithResolvers(r),
		grpc.WithBalancerName(roundrobin.Name),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	go func() {
		time.Sleep(20 * time.Second)
		r.UpdateState(resolver.State{Addresses: []resolver.Address{
			{Addr: "127.0.0.1:5050"},
			{Addr: "127.0.0.1:5051"},
			// {Addr: "127.0.0.1:5052"},
		}})
	}()
}

// PingRouter ...
func PingRouter(r *gin.Engine) {
	r.GET("/ping", getting)
	r.GET("/ping/:name", pingName)
	r.GET("/ping_query", pingQuery)
	r.POST("/ping_json", pingJSON)
	r.GET("/ping_grpc", grpcBalance)
}

// HelloWorld ...
type HelloWorld struct {
	Hello string `form:"hello" json:"hello" xml:"hello"  binding:"required"`
}

func pingJSON(c *gin.Context) {
	var helloPing HelloWorld
	if err := c.ShouldBindJSON(&helloPing); err != nil {
		// c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// if json.User != "manu" || json.Password != "123" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	c.String(http.StatusOK, "Hello %s", helloPing.Hello)
}

func pingQuery(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func pingName(c *gin.Context) {
	var name = c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func getting(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"message": "pong",
	// })

	v := c.Param("k")

	pong := service.PingPong(v)

	c.String(http.StatusOK, pong)
}

func grpcBalance(c *gin.Context) {
	name := c.DefaultQuery("lastname", `defaultname`)

	// defer conn.Close()
	greeter := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := greeter.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// log.Printf("Greeting: %s", res.GetMessage())

	c.String(http.StatusOK, res.GetMessage())
}
