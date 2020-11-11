package controller

import (
	"logical-example/internal/apiserver/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserV1Router ...
func UserV1Router(router *gin.RouterGroup) {
	router.GET("/user", getUserEndpoint)
	router.GET("/user/http", getUserHTTPEndpoint)
}

func getUserEndpoint(c *gin.Context) {

	user, err := service.GetUser("1")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})
}

func getUserHTTPEndpoint(c *gin.Context) {
	result := service.GetUserFromHTTP()
	c.String(http.StatusOK, result)
}
