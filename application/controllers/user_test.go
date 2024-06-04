package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/controllers"
	"github.com/shun198/gin-crm/serializers"
	"github.com/stretchr/testify/assert"
)

// https://gin-gonic.com/docs/testing/
func LoginAdmin() *gin.Engine {
	router := SetUp()
	w := httptest.NewRecorder()
	loginReq := controllers.LoginSerializer{
		EmployeeNumber: "00000001",
		Password:       "test",
	}
	jsonValue, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/api/admin/users/login", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	return router
}

func LoginGeneral() *gin.Engine {
	router := SetUp()
	w := httptest.NewRecorder()
	loginReq := controllers.LoginSerializer{
		EmployeeNumber: "00000001",
		Password:       "test",
	}
	jsonValue, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/api/admin/users/login", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	return router
}

func TestGetAllUsers(t *testing.T) {
	router := LoginAdmin()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/users", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllUsersWithoutLogin(t *testing.T) {
	router := SetUp()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/admin/users", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error": "トークンが必須です"}`, w.Body.String())
}

func TestSendResetPasswordMail(t *testing.T) {
	router := SetUp()
	w := httptest.NewRecorder()
	sendResetPasswordReq := serializers.SendResetPasswordEmailSerializer{
		Email: "test01@example.com",
	}
	jsonValue, _ := json.Marshal(sendResetPasswordReq)
	req, _ := http.NewRequest("POST", "/api/admin/users/send_reset_password_email", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestChangePasswordWithoutLogin(t *testing.T) {
	router := SetUp()
	w := httptest.NewRecorder()
	changePasswordReq := serializers.ChangePasswordSerializer{
		CurrentPassword: "test",
		NewPassword:     "new_pass",
		ConfirmPassword: "new_pass",
	}
	jsonValue, _ := json.Marshal(changePasswordReq)
	req, _ := http.NewRequest("POST", "/api/admin/users/change_password", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
