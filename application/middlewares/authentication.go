package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/prisma/db"
	"github.com/shun198/gin-crm/services"
)

func AuthenticationMiddleware(client *db.PrismaClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
			return
		}
		// ユーザーIDからユーザー情報を取得
		user, err := services.GetUniqueUserByID(userID.(string), client)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "ユーザー情報の取得に失敗しました"})
			return
		}
		// コンテキストにユーザー情報をセット
		c.Set("user", user)
		c.Next()
	}
}

func Authorization(client *db.PrismaClient, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// コンテキストからユーザー情報を取得
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "認証が必要です"})
			return
		}
		auth_user, ok := user.(*db.UserModel)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "ユーザー情報の取得に失敗しました"})
			return
		}
		// ユーザーのロールと要求されたロールを比較してアクセス権を確認
		if string(auth_user.Role) != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "アクセス権がありません"})
			return
		}
		c.Next()
	}
}
