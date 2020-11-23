package controller

import (
	"fmt"
	"log"
	"logical-example/internal/apiserver/service"
	"logical-example/internal/repository"
	"net/http"
	"path/filepath"

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
	router.POST("/user/upload", userUpload)
}

func getUserEndpoint(c *gin.Context) {

	userServiceInstance := c.Value("UserInstance").(service.UserService)

	user, err := userServiceInstance.GetUser("1", repository.UserRepo)

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

func userUpload(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Join("tmp", filepath.Base(file.Filename))
	log.Println(filename)

	// Upload the file to specific dst.
	// c.SaveUploadedFile(file, "tmp")
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
