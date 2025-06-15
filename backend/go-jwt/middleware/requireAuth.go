package middleware

import (
	"go-jwt/initializers"
	"go-jwt/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RequireAuth(c *gin.Context) {
	// Get token from cookie
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		var err error
		tokenString, err = c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Check expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Get user ID (subject) from claims
		userIDStr, ok := claims["sub"].(string)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Convert to UUID
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find user
		var user models.User
		initializers.DB.First(&user, "id = ?", userID)
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach user and userID to context
		c.Set("user", user)
		c.Set("userID", user.ID)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
