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
		userRoutes.POST("/send_reset_password_email", func(c *gin.Context) {
			controllers.SendResetPasswordEmail(c, client)
		})
	}
	return router
}
