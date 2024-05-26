package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/emails"
	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/services"
)

func GetAllUsers(c *gin.Context, client *db.PrismaClient) {
	users, err := services.GetAllUsers(client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func ChangeUserDetails(c *gin.Context, client *db.PrismaClient) {
	userID := c.Param("id")
	user, err := services.GetUniqueUserByID(userID, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "該当するユーザが存在しません"})
		return
	}
	updated_user, err := services.ChangeUserDetails(user, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated_user)
}

func ToggleUserActive(c *gin.Context, client *db.PrismaClient) {
	userID := c.Param("id")
	user, err := services.GetUniqueUserByID(userID, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "該当するユーザが存在しません"})
		return
	}
	toggled_user, err := services.ToggleUserActive(user, client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"is_active": toggled_user.IsActive})
}

func SendInviteUserEmail(c *gin.Context, client *db.PrismaClient) {
	var subject = "ようこそ"
	emails.SendEmail(subject)
}

func ReSendInviteUserEmail(c *gin.Context, client *db.PrismaClient) {
	var subject = "ようこそ"
	emails.SendEmail(subject)
}

func SendResetPasswordEmail(c *gin.Context, client *db.PrismaClient) {
	var subject = "パスワードの再設定"
	emails.SendEmail(subject)
}

func UserInfo(c *gin.Context, client *db.PrismaClient) {
	services.UserInfo()
}
