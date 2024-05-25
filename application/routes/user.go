package routes

import (
	"github.com/gin-gonic/gin"
)

func GetUserRoutes(router *gin.Engine) *gin.Engine {
	userRoutes := router.Group("/api/admin")
	{
		userRoutes.GET("/users")
	}
	return router
}
