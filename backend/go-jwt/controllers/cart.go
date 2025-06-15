package controllers

import (
	"go-jwt/initializers"
	"go-jwt/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddToCart(c *gin.Context) {
	var body struct {
		CropID   uuid.UUID `json:"crop_id"`
		Quantity int       `json:"quantity"`
	}
	if err := c.BindJSON(&body); err != nil || body.Quantity <= 0 {
		c.JSON(400, gin.H{
			"error": "Invalid input",
		})
	}
	userID, exits := c.Get("userID")
	if !exits {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	cartItem := models.CartItem{
		UserID:   userID.(uint),
		CropID:   body.CropID,
		Quantity: body.Quantity,
	}
	if err := initializers.DB.Create(&cartItem).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not add to cart"})
		return
	}
	c.JSON(200, gin.H{"message": "Item added to cart"})
}
