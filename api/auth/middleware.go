package auth

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (s *authService) RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header)
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		slog.Error("error while parsing token", "error", err, "token", token)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		slog.Error("error while parsing claims", "error", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	exp, ok := claims["exp"].(float64)
	if !ok || float64(time.Now().Unix()) > exp {
		slog.Error("token expired", "error", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userIDFloat, ok := claims["sub"].(float64) // Extract userID as float64
	if !ok {
		slog.Error("error while parsing user id", "userID", userIDFloat)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userID := uint(userIDFloat)

	user, err := s.userService.GetUserByID(c, userID)
	if err != nil {
		slog.Error("error while getting user by id", "error", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}
	c.Set("user", user)

	c.Next()
}
