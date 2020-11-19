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

func TestPing(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Params = []gin.Param{{Key: "k", Value: "v"}}

	getting(c)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())

	// if w.Code != 200 {
	// 	b, _ := ioutil.ReadAll(w.Body)
	// 	t.Error(w.Code, string(b))
	// }
}
