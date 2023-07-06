package main

import (
	"github.com/charliscript/go-crud/initializers"
	"github.com/charliscript/go-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.Users{})
}
