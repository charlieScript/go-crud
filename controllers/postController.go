package controllers

import (
	"log"
	"net/http"

	"github.com/charliscript/go-crud/initializers"
	"github.com/charliscript/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var Body struct {
		Body  string `binding:"required"`
		Title string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: Body.Title, Body: Body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "An error occurred",
		})
	}

	c.JSON(200, post)
}

func PostsGet(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, posts)
}

func PostsGetById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := c.ShouldBindUri(&post); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}

	result := initializers.DB.First(&post, id)

	log.Fatal(result.Error.Error())

	c.JSON(200, post)
}
