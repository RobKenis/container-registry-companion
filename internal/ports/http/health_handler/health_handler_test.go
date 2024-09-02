package health_handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/robkenis/container-registry-companion/internal/ports/http/health_handler"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := health_handler.Handler{}
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, rr.Body.String(), `{"status":"UP"}`+"\n")
}
