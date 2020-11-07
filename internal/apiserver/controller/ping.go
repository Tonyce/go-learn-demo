package controller

import (
	"logical-example/internal/apiserver/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingRouter ...
func PingRouter(r *gin.Engine) {
	r.GET("/ping", getting)
	r.GET("/ping/:name", pingName)
	r.GET("/ping_query", pingQuery)
	r.POST("/ping_json", pingJSON)
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
