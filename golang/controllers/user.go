package controllers

import (
	"leiloa/db"
	"leiloa/helpers"
	"leiloa/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		State    string `json:"state"`
		Cpf      string `json:"cpf"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	if cpfIsValid := helpers.CpfIsValid(body.Cpf); !cpfIsValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cpf is not valid"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to hash the Password"})
		log.Println("Failed to hash Password", err)
		return
	}

	var newUser models.User

	newUser.Email = body.Email
	newUser.Cpf = body.Cpf
	newUser.Name = body.Name
	newUser.Password = body.Password
	newUser.State = body.State

	body.Password = string(hash)

	if err := db.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		log.Println("Failed to create user", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User Created Successfully"})
}
