package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCommonRoutes(router *gin.Engine) *gin.Engine {
	commonRoutes := router.Group("/api")
	{
		commonRoutes.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "pass",
			})
		})
	}
	return router
}
