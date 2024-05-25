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
		userRoutes.GET("", func(c *gin.Context) {
			controllers.GetUsers(c, client)
		})
		userRoutes.GET("/get_csrf_token", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"csrf-token": csrf.GetToken(c),
			})
		})
	}
	return router
}
