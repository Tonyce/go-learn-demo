package testapiserver

import (
	"logical-example/internal/apiserver"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/jarcoal/httpmock"
)

func TestUserHTTPRoute(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://127.0.0.1:8989",
		httpmock.NewStringResponder(200, `http mock`))

	router := apiserver.SetupRouter()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/v1/user/http", nil)
	req.Header.Set("User-Agent", "http-test")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "http mock", w.Body.String())
}
