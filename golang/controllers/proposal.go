package controllers

import (
	"leiloa/db"
	"leiloa/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProposal(c *gin.Context) {
	userData, _ := c.Get("user")

	user := userData.(models.User)

	var body struct {
		Value     float32 `json:"value"`
		AuctionId string  `json:"auctionId"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	var newProposal models.Proposal

	newProposal.AuctionId = body.AuctionId
	newProposal.Value = body.Value
	newProposal.User = user

	if err := db.DB.Create(&newProposal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create proposal"})
		log.Println("Failed to create proposal", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Proposal Created Successfully"})
}
