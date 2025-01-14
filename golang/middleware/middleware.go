package middleware

import (
	"fmt"
	"leiloa/db"
	"leiloa/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Status": "Not Logged In"})
		c.Abort()
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token has expired"})
			c.Abort()
		}

		var user models.User

		user.Id = claims["sub"].(string)

		db.DB.First(&user)

		if user.Id == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "user not found"})
			c.Abort()
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "not logged in"})
		c.Abort()
	}
}
