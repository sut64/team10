package controller

import (
	"net/http"
	"github.com/sut64/team10/entity"

	"github.com/gin-gonic/gin"
)

func ListBloodType(c *gin.Context) {
	var bloodtype []entity.BloodType
	if err := entity.DB().Raw("SELECT * FROM blood_types").Scan(&bloodtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bloodtype})
}

func GetBloodType(c *gin.Context) {
	var bloodtype entity.BloodType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM blood_types WHERE id = ?", id).Scan(&bloodtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bloodtype})
}

func CreateBloodType(c *gin.Context) {
	var bloodtype entity.BloodType
	if err := c.ShouldBindJSON(&bloodtype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&bloodtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bloodtype})
}

func UpdateBloodType(c *gin.Context) {
	var bloodtype entity.BloodType
	if err := c.ShouldBindJSON(&bloodtype); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bloodtype.ID).First(&bloodtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bloodtype not found"})
		return
	}

	if err := entity.DB().Save(&bloodtype).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bloodtype})
}

func DeleteBloodType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM blood_types WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bloodtype not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
