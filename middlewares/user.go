package middlewares

import (
	"context"
	"net/http"
	"strings"
	"task/config"
	"task/helpers"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Token is required")
			return
		}
		var tokenString string

		splitToken := strings.Split(tokenStr, "Bearer ")
		if len(splitToken) > 1 {
			tokenString = splitToken[1]
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid token")
			return
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&helpers.JWTClaimForUser{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(helpers.JwtKey), nil
			},
		)
		if err != nil {
			c.AbortWithStatusJSON(403, gin.H{"message": err.Error()})
			return
		}
		claims, ok := token.Claims.(*helpers.JWTClaimForUser)
		if !ok {
			c.AbortWithStatusJSON(400, gin.H{"message": "couldn't parse claims"})
			return
		}
		if claims.ExpiresAt < time.Now().Local().Unix() {
			c.AbortWithStatusJSON(403, gin.H{"message": "token expired"})
			return
		}

		db, err := config.ConnDB()
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
			return
		}
		defer db.Close()

		var user_id string
		if err := db.QueryRow(context.Background(), "SELECT id FROM users WHERE id = $1 AND deleted_at IS NULL", claims.UserID).Scan(&user_id); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}

}
