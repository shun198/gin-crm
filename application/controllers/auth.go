package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/config"
	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/serializers"
	"github.com/shun198/gin-crm/services"
)

type Claims struct {
	EmployeeNumber string `json:"employee_number"`
	jwt.StandardClaims
}

func Login(c *gin.Context, client *db.PrismaClient) {
	var req serializers.LoginSerializer
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "社員番号もしくはパスワードが間違っています"})
		return
	}
	user, err := services.GetUniqueUserByEmployeeNumber(req.EmployeeNumber, client)
	if err != nil {
		// 実行時間から社員番号が正しいかわからないよう対策する
		config.HashPassword(req.Password)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "社員番号もしくはパスワードが間違っています"})
		return
	}
	check := config.CheckPasswordHash(user.Password, req.Password)
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "社員番号もしくはパスワードが間違っています"})
		return
	}
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		EmployeeNumber: user.EmployeeNumber,
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("ISS"),
			Audience:  os.Getenv("AUD"),
			Subject:   os.Getenv("SUB"),
			ExpiresAt: expirationTime.Unix(),
		},
	}
	var jwtKey = []byte(os.Getenv("SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "トークンの生成に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Logout(c *gin.Context) {
	// services.Logout()
	c.JSON(http.StatusOK, gin.H{})
}
