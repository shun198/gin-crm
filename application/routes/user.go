package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/controllers"
	"github.com/shun198/gin-crm/prisma/db"
	csrf "github.com/utrack/gin-csrf"
)

func GetUserRoutes(router *gin.Engine, client *db.PrismaClient) *gin.Engine {
	userRoutes := router.Group("/api/admin/users")
	{
		userRoutes.GET("/get_csrf_token", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"csrf-token": csrf.GetToken(c),
			})
		})
		userRoutes.POST("/login", func(c *gin.Context) {
			controllers.Login(c, client)
		})
		userRoutes.POST("/logout", func(c *gin.Context) {
			controllers.Logout(c)
		})
		userRoutes.GET("", func(c *gin.Context) {
			controllers.GetAllUsers(c, client)
		})
		userRoutes.PATCH("/:id/change_user_details", func(c *gin.Context) {
			controllers.ChangeUserDetails(c, client)
		})
		userRoutes.POST("/:id/toggle_user_active", func(c *gin.Context) {
			controllers.ToggleUserActive(c, client)
		})
		userRoutes.POST("/send_invite_user_email", func(c *gin.Context) {
			controllers.SendInviteUserEmail(c, client)
		})
		userRoutes.POST(":id/resend_invitation", func(c *gin.Context) {
			controllers.ReSendInviteUserEmail(c, client)
		})
		userRoutes.POST("/send_reset_password_email", func(c *gin.Context) {
			controllers.SendResetPasswordEmail(c, client)
		})
		userRoutes.POST("/verify_user", func(c *gin.Context) {
			controllers.SendInviteUserEmail(c, client)
		})
		userRoutes.POST("/change_password", func(c *gin.Context) {
			controllers.ChangePassword(c, client)
		})
		userRoutes.POST("/reset_password", func(c *gin.Context) {
			controllers.ResetPassword(c, client)
		})
		userRoutes.POST("/check_invitation_token", func(c *gin.Context) {
			controllers.CheckInvitationToken(c, client)
		})
		userRoutes.POST("/check_reset_password_token", func(c *gin.Context) {
			controllers.CheckResetPasswordToken(c, client)
		})
		userRoutes.POST("/user_info", func(c *gin.Context) {
			controllers.UserInfo(c, client)
		})
	}
	return router
}
