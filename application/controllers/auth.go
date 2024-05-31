package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/config"
	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/serializers"
	"github.com/shun198/gin-crm/services"
)

func Login(c *gin.Context, client *db.PrismaClient) {
	var req serializers.LoginSerializer
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "社員番号もしくはパスワードが間違っています"})
		return
	}
	user, err := services.GetUniqueUserByEmployeeNumber(req.EmployeeNumber, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "社員番号もしくはパスワードが間違っています"})
		return
	}
	check := config.CheckPasswordHash(user.Password, req.Password)
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "社員番号もしくはパスワードが間違っています"})
		return
	}
	services.Login()
}

func Logout(c *gin.Context) {
	// services.Logout()
	c.JSON(http.StatusOK, gin.H{})
}
