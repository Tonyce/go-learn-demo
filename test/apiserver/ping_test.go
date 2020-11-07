package testapiserver

import (
	"bytes"
	"logical-example/internal/apiserver"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func init() {
	// gin.SetMode(gin.TestMode)
}

func TestPingRoute(t *testing.T) {
	t.Parallel()
	router := apiserver.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)
	req.Header.Set("User-Agent", "http-test")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, gin.H{
	// 	"message": "pong",
	// }, w.Body.String())
}

func TestPingNameRoute(t *testing.T) {
	t.Parallel()
	router := apiserver.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping/thegin", nil)
	req.Header.Set("User-Agent", "http-test")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello thegin", w.Body.String())
}

func TestPingQueryRoute(t *testing.T) {
	t.Parallel()
	router := apiserver.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping_query?firstname=Jane&lastname=Doe", nil)
	req.Header.Set("User-Agent", "http-test")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello Jane Doe", w.Body.String())
}

func TestPingJsonRoute(t *testing.T) {
	t.Parallel()
	router := apiserver.SetupRouter()

	w := httptest.NewRecorder()
	var jsonStr = []byte(`{hello":"world"}`)
	req, _ := http.NewRequest("POST", "/ping_json?firstname=Jane&lastname=Doe", bytes.NewBuffer(jsonStr))
	req.Header.Set("User-Agent", "http-test")
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("X-Custom-Header", "myvalue")

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	// assert.Equal(t, "Hello world", w.Body.String())
}

/**
package main

import (
   "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
)

funcTestPingRoute(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET","/ping", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Equal(t,"pong", w.Body.String())
}
*/
