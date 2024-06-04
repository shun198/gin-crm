package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/config"
	"github.com/shun198/gin-crm/controllers"
	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/routes"
	"github.com/stretchr/testify/assert"
)

func SetUp() (*gin.Engine, *db.PrismaClient) {
	r := gin.Default()
	client, err := config.StartDatabase()
	if err != nil {
		panic(err)
	}
	router := routes.GetUserRoutes(r, client)
	return router, client
}

func TestLogin(t *testing.T) {
	router, _ := SetUp()
	w := httptest.NewRecorder()
	loginReq := controllers.LoginSerializer{
		EmployeeNumber: "00000001",
		Password:       "test",
	}
	jsonValue, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/api/admin/users/login", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginWithWrongPassword(t *testing.T) {
	router, _ := SetUp()
	w := httptest.NewRecorder()
	loginReq := controllers.LoginSerializer{
		EmployeeNumber: "00000001",
		Password:       "wrong_password",
	}
	jsonValue, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/api/admin/users/login", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"msg": "社員番号もしくはパスワードが間違っています"}`, w.Body.String())
}

func TestLoginWithWrongEmployeeNumber(t *testing.T) {
	router, _ := SetUp()
	w := httptest.NewRecorder()
	loginReq := controllers.LoginSerializer{
		EmployeeNumber: "99999999",
		Password:       "test",
	}
	jsonValue, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/api/admin/users/login", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"msg": "社員番号もしくはパスワードが間違っています"}`, w.Body.String())
}
