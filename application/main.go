package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	database "github.com/shun198/gin-crm/config"
	docs "github.com/shun198/gin-crm/docs"
	"github.com/shun198/gin-crm/middlewares"
	"github.com/shun198/gin-crm/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	csrf "github.com/utrack/gin-csrf"
)

// @title gin-crm API
// @version 1.0
// @description This is a sample crm project
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	r.Use(middlewares.LoggerMiddleWare())
	r.Use(gin.Recovery())
	// https://pkg.go.dev/github.com/marktohark/gin-csrf#section-readme
	store := cookie.NewStore([]byte("cookie_secret"))
	r.Use(sessions.Sessions("session_name_in_cookie", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "csrf_token",
		ErrorFunc: func(c *gin.Context) {
			c.String(http.StatusForbidden, "無効なCSRFトークンです")
			c.Abort()
		},
	}))
	err := database.StartDatabase()
	if err != nil {
		panic(err)
	}
	// @BasePath /api

	// @Summary ping example
	// @Schemes
	// @Description do ping
	// @Tags example
	// @Accept json
	// @Produce json
	// @Success 200 {string} Helloworld
	// @Router /api/health [get]
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pass",
		})
	})
	r.GET("/api/admin/users/get_csrf_token", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"csrf-token": csrf.GetToken(c),
		})
	})
	routes.GetUserRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8000")
}
