package controller

import (
	"logical-example/internal/apiserver/service"
	"logical-example/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("UserInstance", service.UserServiceInstance)
		c.Next()
	}
}

// UserV1Router ...
func UserV1Router(router *gin.RouterGroup) {
	router.GET("/user", userMiddleware(), getUserEndpoint)
	router.GET("/user/http", userMiddleware(), getUserHTTPEndpoint)
}

// TODO context 设置 serviceInstance Repository
func getUserEndpoint(c *gin.Context) {

	userServiceInstance := c.Value("UserInstance").(service.UserService)

	user, err := userServiceInstance.GetUser("1", model.MongoUserRepository)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})
}

func getUserHTTPEndpoint(c *gin.Context) {
	userServiceInstance := c.Value("UserInstance").(service.UserService)
	result := userServiceInstance.GetUserFromHTTP()
	c.String(http.StatusOK, result)
}
