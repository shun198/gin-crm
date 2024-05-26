package main

import (
	"database/sql"
	"log"
	"os"

	// "net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
	database "github.com/shun198/gin-crm/config"
	_ "github.com/shun198/gin-crm/docs"
	"github.com/shun198/gin-crm/middlewares"
	"github.com/shun198/gin-crm/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// csrf "github.com/utrack/gin-csrf"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.LoggerMiddleWare())
	r.Use(gin.Recovery())
	// https://github.com/gin-contrib/cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
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
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@db:5432/postgres")
	if err != nil {
		log.Fatal("データベースの接続の切断に失敗しました:%v", err)
	}
	store, err := postgres.NewStore(db, []byte(os.Getenv("SECRET_KEY")))
	r.Use(sessions.Sessions("session", store))
	// https://pkg.go.dev/github.com/marktohark/gin-csrf#section-readme
	// r.Use(csrf.Middleware(csrf.Options{
	// 	Secret: "csrf_token",
	// 	ErrorFunc: func(c *gin.Context) {
	// 		c.String(http.StatusForbidden, "無効なCSRFトークンです")
	// 		c.Abort()
	// 	},
	// }))
	routes.GetCommonRoutes(r)
	routes.GetUserRoutes(r, client)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8000")
}
