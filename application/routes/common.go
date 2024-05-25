package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCommonRoutes(router *gin.Engine) *gin.Engine {
	commonRoutes := router.Group("/api")
	{
		// HealthCheck godoc
		// @Summary Health Check
		// @Description Check if the server is running
		// @Tags health
		// @Accept json
		// @Produce json
		// @Success 200 {object} map[string]string{"msg":"pass"}
		// @Router /health [get]
		commonRoutes.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "pass",
			})
		})
	}
	return router
}
