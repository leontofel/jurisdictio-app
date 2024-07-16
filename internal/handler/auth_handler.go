package handler

import (
	"justice-app/internal/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	type Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var login Login
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid format")
		return
	}

	var user model.User
	if err := db.Where("email = ?", login.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, "User not found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, "Incorrect password")
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	// Claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID       // User ID
	claims["email"] = user.Email // User email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate token
	secret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not generate token")
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Register(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	type RegistrationInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	var input RegistrationInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid input format")
		return
	}

	// Check if user already exists
	var existingUser model.User
	if err := db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, "User already exists")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Could not hash password")
		return
	}

	// Create new user
	newUser := model.User{
		Email:    input.Email,
		Password: string(hashedPassword),
		Name:     input.Name,
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Could not save user")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
