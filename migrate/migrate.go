package main

import (
	"GoLang/initializers"
	"GoLang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	initializers.Database.AutoMigrate(&models.User{})
}
