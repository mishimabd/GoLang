package controllers

import (
	"GoLang/initializers"
	"GoLang/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateUser(c *gin.Context) {
	//create struct for users
	//password will be hashed by JWT
	var body struct {
		Name     string
		Surname  string
		Email    string
		Password string
		Role     bool
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

	post := models.User{Name: body.Name, Surname: body.Surname, Email: body.Email, Password: string(hash), Role: body.Role}
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

func DeleteUser(c *gin.Context) {

	//getting id of from parametr
	id := c.Param("id")

	//delete user
	initializers.Database.Delete(&models.User{}, id)

	//respond
	c.Status(200)
}

func GetAllUsers(c *gin.Context) {
	//create object of users
	var users []models.User

	//find from database
	initializers.Database.Find(&users)

	//respond from postman
	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetOneUser(c *gin.Context) {
	//getting one id from paran
	id := c.Param("id")
	//creating object
	var user models.User
	//initialize bd for users
	initializers.Database.First(&user, id)
	//respond

	if user.ID == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "there is no user with id like this",
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"users": user,
		})
	}
}

func UpdateOneUser(c *gin.Context) {
	//the id of the url
	id := c.Param("id")

	//get the data of request body
	var body struct {
		Name    string
		Surname string
	}

	c.Bind(&body)
	//find the post we are updating
	var user models.User
	initializers.Database.First(&user, id)

	//update it
	initializers.Database.Model(&user).Updates(models.User{Name: body.Name, Surname: body.Surname})

	//respond

	c.JSON(200, gin.H{
		"post": user,
	})
}
