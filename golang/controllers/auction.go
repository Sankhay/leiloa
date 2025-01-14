package controllers

import (
	"leiloa/db"
	"leiloa/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAuction(c *gin.Context) {
	userData, _ := c.Get("user")

	user := userData.(models.User)

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	name := form.Value["name"][0]
	description := form.Value["description"][0]
	category := form.Value["category"][0]

	var newAuction models.Auction

	newAuction.Name = name
	newAuction.Description = description
	newAuction.CategoryId = category
	newAuction.Owner = user

	filesDirectory := "files/"

	filesPath := []string{}

	for _, file := range files {

		fileName := uuid.New().String()
		filePath := filesDirectory + fileName
		filesPath = append(filesPath, filePath)

		c.SaveUploadedFile(file, filePath)
	}

	if err := db.DB.Create(&newAuction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating new auction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Proposal Created Successfully"})
}
