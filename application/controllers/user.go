package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/services"
)

func GetUsers(c *gin.Context, client *db.PrismaClient) {
	users, err := services.GetAllUsers(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func ChangeUserDetails(c *gin.Context, client *db.PrismaClient) {
	user, err := services.ChangeUserDetails(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
