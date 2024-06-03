package middlewares

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/shun198/gin-crm/prisma/db"
)

type Claims struct {
	UserID int     `json:"user_id"`
	Role   db.Role `json:"role"`
	jwt.StandardClaims
}

func AuthenticationMiddleware() gin.HandlerFunc {
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
		c.Set("role", claims.Role)
		c.Set("tokenString", tokenString)
		c.Next()
	}
}

func AuthorizationMiddleware(roles ...db.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role not found"})
			c.Abort()
			return
		}

		roleEnum, ok := role.(db.Role)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid role type"})
			c.Abort()
			return
		}

		authorized := false
		for _, r := range roles {
			if roleEnum == r {
				authorized = true
				break
			}
		}

		if !authorized {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access forbidden: insufficient role"})
			c.Abort()
			return
		}

		c.Next()
	}
}
