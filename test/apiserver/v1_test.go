package testapiserver

import (
	"logical-example/internal/apiserver"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func init() {
	// gin.SetMode(gin.TestMode)

}

func TestV1Route(t *testing.T) {
	router := apiserver.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/login", nil)
	req.Header.Set("User-Agent", "http-test")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
