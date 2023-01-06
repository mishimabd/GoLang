package controllers

import (
	"GoLang/initializers"
	"GoLang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func SignUp(c *gin.Context) {
	//create struct for users
	//password will be hashed by JWT
	var body struct {
		Name     string
		Surname  string
		Email    string
		Password string
	}
	//scanning body struct
	c.Bind(&body)
	//hashing your password by JWT token
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	post := models.User{Name: body.Name, Surname: body.Surname, Email: body.Email, Password: string(hash), Role: false}
	//Creating user
	result := initializers.Database.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	//Respond
	c.JSON(200, gin.H{
		"user": post})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	//scanning body struct
	c.Bind(&body)
	//hashing your password by JWT token
	var user models.User

	initializers.Database.First(&user, "email=?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//generating jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid to create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GetCourses(c *gin.Context) {
	//create object of users
	var course []models.Course

	//find from database
	initializers.Database.Find(&course)

	//respond from postman
	c.JSON(200, gin.H{
		"courses": course,
	})
}
