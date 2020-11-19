package controller

import (
	"logical-example/internal/apiserver/service"
	"logical-example/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserV1Router ...
func UserV1Router(router *gin.RouterGroup) {
	router.GET("/user", getUserEndpoint)
	router.GET("/user/http", getUserHTTPEndpoint)
}

// TODO context 设置 serviceInstance Repository
func getUserEndpoint(c *gin.Context) {

	user, err := service.UserServiceInstance.GetUser("1", model.MongoUserRepository)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})
}

func getUserHTTPEndpoint(c *gin.Context) {
	result := service.UserServiceInstance.GetUserFromHTTP()
	c.String(http.StatusOK, result)
}
