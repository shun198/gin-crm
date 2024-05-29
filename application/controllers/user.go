package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shun198/gin-crm/emails"
	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/serializers"
	"github.com/shun198/gin-crm/services"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

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
	var req serializers.ChangeUserDetailsSerializer
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なリクエストです"})
		return
	}
	err := validate.Struct(req)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			// カスタムエラーメッセージを生成
			var errMsg string
			switch err.Tag() {
			case "required":
				errMsg = fmt.Sprintf("%sは必須項目です", err.Field())
			case "max":
				errMsg = fmt.Sprintf("%sが長すぎます", err.Field())
			case "email":
				errMsg = fmt.Sprintf("%sは有効なメールアドレスでなければなりません", err.Field())
			case "oneof":
				errMsg = fmt.Sprintf("%sは%sのいずれかでなければなりません", err.Field(), err.Param())
			default:
				errMsg = fmt.Sprintf("%sは無効です", err.Field())
			}
			validationErrors = append(validationErrors, errMsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}
	_, err = services.GetUniqueUserByEmail(*req.Email, client)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "登録されていないメールアドレスを入力してください"})
		return
	}
	services.ChangeUserDetails(req, userID, client)
	c.JSON(http.StatusOK, err)
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
	var req serializers.SendInviteUserEmailSerializer
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なリクエストです"})
		return
	}
	err := validate.Struct(req)
	if err != nil {
		var validationErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			// カスタムエラーメッセージを生成
			var errMsg string
			switch err.Tag() {
			case "required":
				errMsg = fmt.Sprintf("%sは必須項目です", err.Field())
			case "max":
				errMsg = fmt.Sprintf("%sが長すぎます", err.Field())
			case "email":
				errMsg = fmt.Sprintf("%sは有効なメールアドレスでなければなりません", err.Field())
			case "employee_number":
				errMsg = "数字8桁の社員番号を入力してください"
			case "oneof":
				errMsg = fmt.Sprintf("%sは%sのいずれかでなければなりません", err.Field(), err.Param())
			default:
				errMsg = fmt.Sprintf("%sは無効です", err.Field())
			}
			validationErrors = append(validationErrors, errMsg)
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}
	_, err = services.GetUniqueUserByEmail(*req.Email, client)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "登録されていないメールアドレスを入力してください"})
		return
	}
	_, err = services.GetUniqueUserByEmployee_number(*req.EmployeeNumber, client)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "登録されていない社員番号を入力してください"})
		return
	}
	var subject = "ようこそ"
	emails.SendEmail(subject)
}

func ReSendInviteUserEmail(c *gin.Context, client *db.PrismaClient) {
	userID := c.Param("id")
	// Invitationトークンから該当するユーザがないか探す
	user, err := services.GetUniqueUserByID(userID, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "該当するユーザが存在しません"})
		return
	}
	// ユーザが無効化されているかどうか
	if !user.IsActive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ユーザが無効化されているため招待メールを再送信できません"})
		return
	}
	if user.IsVerified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ユーザの認証がすでに完了しているため招待メールを再送信できません"})
		return
	}
	var subject = "ようこそ"
	emails.SendEmail(subject)
}

func SendResetPasswordEmail(c *gin.Context, client *db.PrismaClient) {
	var subject = "パスワードの再設定"
	emails.SendEmail(subject)
}

func VerifyUser(c *gin.Context, client *db.PrismaClient) {
	services.VerifyUser()
}

func ChangePassword(c *gin.Context, client *db.PrismaClient) {
	services.ChangePassword()
}

func ResetPassword(c *gin.Context, client *db.PrismaClient) {
	services.ResetPassword()
}

func CheckInvitationToken(c *gin.Context, client *db.PrismaClient) {
	services.CheckInvitationToken()
}

func CheckResetPasswordToken(c *gin.Context, client *db.PrismaClient) {
	services.CheckResetPasswordToken()
}

func UserInfo(c *gin.Context, client *db.PrismaClient) {
	services.UserInfo()
}
