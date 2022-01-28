package controller

import (
	"net/http"
	"github.com/sut64/team10/entity"

	"github.com/gin-gonic/gin"
)

func CreateMedicine(c *gin.Context) {
	var medicne entity.Medicine
	if err := c.ShouldBindJSON(&medicne); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&medicne).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": medicne})
}

// GET /medicne/:id
func GetMedicne(c *gin.Context) {
	var medicne entity.Medicine
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medicnes WHERE id = ?", id).Scan(&medicne).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicne})
}

// GET /medicne
func ListMedicine(c *gin.Context) {
	var medicne []entity.Medicine
	if err := entity.DB().Raw("SELECT * FROM medicines").Scan(&medicne).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicne})
}

// DELETE /medicne/:id
func DeleteMedicine(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medicnes WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicne not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /medicne
func UpdateMedicine(c *gin.Context) {
	var medicne entity.Medicine
	if err := c.ShouldBindJSON(&medicne); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", medicne.ID).First(&medicne); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicne not found"})
		return
	}

	if err := entity.DB().Save(&medicne).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicne})
}

