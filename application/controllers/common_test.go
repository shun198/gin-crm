package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/routes"
	"github.com/stretchr/testify/assert"
)

// https://gin-gonic.com/docs/testing/
func TestHealthCheck(t *testing.T) {
	r := gin.Default()
	router := routes.GetCommonRoutes(r)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"msg": "pass"}`, w.Body.String())
}
