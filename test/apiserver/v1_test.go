package testapiserver

import (
	"fmt"
	"logical-example/internal/apiserver"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func init() {
	// gin.SetMode(gin.TestMode)

}

func beforeAll() {
	t := time.Now()
	fmt.Println("before ", t.Format(time.RFC3339))
}

func afterAll() {
	fmt.Println("after ", time.Now())
}

func TestMain(m *testing.M) {
	beforeAll()
	code := m.Run()
	// shutdown()
	afterAll()
	os.Exit(code)
}

func TestV1Route(t *testing.T) {
	router := apiserver.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/login", nil)
	req.Header.Set("User-Agent", "http-test")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestUserV1Route(t *testing.T) {
	router := apiserver.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/user", nil)
	req.Header.Set("User-Agent", "http-test")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"name":"test"}`, w.Body.String())
}
