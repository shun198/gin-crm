package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "無効なトークンです"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]

		claims := &Claims{}
		var jwtKey = []byte(os.Getenv("SECRET_KEY"))
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "無効なトークンです"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("tokenString", tokenString)
		c.Next()
	}
}
