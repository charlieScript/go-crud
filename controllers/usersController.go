package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/charliscript/go-crud/initializers"
	"github.com/charliscript/go-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string `binding:"required"`
		Password string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "An error occurred",
		})
	}
	user := models.Users{Email: body.Email, Password: string(hashPassword)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "An error occurred",
		})
		return
	}

	c.JSON(200, gin.H{"message": "Welcome" + user.Email})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `binding:"required"`
		Password string `binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.Users

	initializers.DB.Find(&user, "email = ?", body.Email)

	if user.Email == "" {
		c.JSON(404, gin.H{
			"message": "Invalid username or password",
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.JSON(404, gin.H{
			"message": "Invalid username or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Println(err)
		log.Println(os.Getenv("JWT_SECRET"))
		c.JSON(400, gin.H{
			"message": "An error occured",
		})
		return
	}

	// cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*24, "", "", false, false)

	c.JSON(200, gin.H{"user": user.Email, "token": tokenString})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user)
}
