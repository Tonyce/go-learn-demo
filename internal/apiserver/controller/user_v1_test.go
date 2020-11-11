package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func init() {
	gin.SetMode(gin.TestMode)
	// common.InitMongoDB()
}

func TestUserV1(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "k", Value: "test"}}

	getUserEndpoint(c)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, `{"name":"test"}`, w.Body.String())
	assert.Equal(t, `{"name":"1"}`, w.Body.String())
}

func TestUserHTTP(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "k", Value: "test"}}

	getUserHTTPEndpoint(c)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, `{"name":"test"}`, w.Body.String())
	assert.Equal(t, `GetUserFromHTTP`, w.Body.String())
}
