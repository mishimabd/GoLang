package middleware

import (
	"GoLang/initializers"
	"GoLang/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		fmt.Println("test 0")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("test 1")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User

		initializers.Database.First(&user, claims["sub"])

		if user.ID == 0 {
			fmt.Println("test 2")
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
