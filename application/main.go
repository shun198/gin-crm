package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	database "github.com/shun198/gin-crm/config"
	"github.com/shun198/gin-crm/middleware"
	"github.com/shun198/gin-crm/router"
	csrf "github.com/utrack/gin-csrf"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleWare())
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
	router.GetUserRoutes(r)
	r.Run(":8000")
}
