package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// V1Router ...
func V1Router(router *gin.Engine) {
	v1 := router.Group("/v1")

	{
		v1.GET("/login", loginEndpoint)
		UserV1Router(v1)
	}
}

func loginEndpoint(c *gin.Context) {
	// ex, exist := c.Get("example")
	// if exist {
	// 	log.Println(ex)
	// }
	c.String(http.StatusOK, "v1 login")
}
