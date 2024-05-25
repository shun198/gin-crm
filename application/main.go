package main

import (
	"log"
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
	client, err := database.StartDatabase()
	if err != nil {
		log.Fatal("データベースとの接続に失敗しました:%v", err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatal("データベースの接続の切断に失敗しました:%v", err)
		}
	}()
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
	routes.GetUserRoutes(r, client)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8000")
}
