package main

import (
	"github.com/charliscript/go-crud/controllers"
	"github.com/charliscript/go-crud/initializers"
	"github.com/charliscript/go-crud/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.PostsCreate)
	r.GET("/", controllers.PostsGet)
	r.GET("/:id", controllers.PostsGetById)
	r.POST("/auth/signup", controllers.Signup)
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/validate", middlewares.RequireAuth, controllers.Validate)
	r.Run()
}
