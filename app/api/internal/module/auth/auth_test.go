package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLoginEndpointRejectsInvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/api/auth/login", LoginHandler(nil, nil))

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/auth/login",
		strings.NewReader(`{"email":"not-an-email","password":"short"}`),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d; body: %s",
			rec.Code, http.StatusBadRequest, rec.Body.String())
	}
}
