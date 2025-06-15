package controllers

import (
	"go-jwt/initializers"
	"go-jwt/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CropsCreate(c *gin.Context) {
	//Get data of req body
	var crop models.Crop

	if err := c.ShouldBindJSON(&crop); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid crop data",
		})
		return
	}
	if crop.Name == "" || crop.Category == "" || crop.Price <= 0 || crop.Quantity <= 0 || crop.FarmerID == uuid.Nil {
		c.JSON(400, gin.H{"error": "Missing or invalid required fields"})
		return
	}

	if err := initializers.DB.Create(&crop).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create crop"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Crop created successfully",
		"crop":    crop,
	})
}

func Cropindex(c *gin.Context) {
	var crops []models.Crop
	result := initializers.DB.Find(&crops)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"crops": crops,
	})
}

func CropsShow(c *gin.Context) {
	idParam := c.Param("id")

	// Convert string ID to UUID
	uid, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID format"})
		return
	}

	var crop models.Crop
	result := initializers.DB.First(&crop, "id = ?", uid)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Crop not found"})
		return
	}

	c.JSON(200, gin.H{
		"crop": crop,
	})
}

func CropsUpdate(c *gin.Context) {
	idParam := c.Param("id")
	uid, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID format"})
		return
	}

	var existingCrop models.Crop
	if err := initializers.DB.First(&existingCrop, "id = ?", uid).Error; err != nil {
		c.JSON(404, gin.H{"error": "Crop not found"})
		return
	}

	var updateData models.Crop
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Update fields (only the updatable ones)
	existingCrop.Name = updateData.Name
	existingCrop.Category = updateData.Category
	existingCrop.Price = updateData.Price
	existingCrop.Quantity = updateData.Quantity
	existingCrop.ImageURL = updateData.ImageURL

	if err := initializers.DB.Save(&existingCrop).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update crop"})
		return
	}

	c.JSON(200, gin.H{"message": "Crop updated", "crop": existingCrop})
}

func CropsDelete(c *gin.Context) {
	idParam := c.Param("id")
	uid, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid UUID format"})
		return
	}

	if err := initializers.DB.Delete(&models.Crop{}, "id = ?", uid).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete crop"})
		return
	}

	c.JSON(200, gin.H{"message": "Crop deleted"})
}

func MyCrops(c *gin.Context) {
	// Get user ID from context
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Fetch crops for this user
	var crops []models.Crop
	if err := initializers.DB.Where("farmer_id = ?", userID).Find(&crops).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch crops"})
		return
	}

	c.JSON(200, gin.H{"crops": crops})
}
