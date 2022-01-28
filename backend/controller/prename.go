package controller

import (
	"net/http"


	"github.com/gin-gonic/gin"
	"github.com/sut64/team10/entity"
)

// POST /PrenameController
func CreatePrename(c *gin.Context) {
	var prenames entity.Prename
	if err := c.ShouldBindJSON(&prenames); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&prenames).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": prenames})
}

// GET /prename/:id
func GetPrename(c *gin.Context) {
	var prenames entity.Prename
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM prenames WHERE id = ?", id).Find(&prenames).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prenames})
}

// GET /return_ways
func ListPrenames(c *gin.Context) {
	var prenames []entity.Prename
	if err := entity.DB().Raw("SELECT * FROM prenames").Find(&prenames).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prenames})
}

// DELETE /return_ways/:id
func DeletePrename(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM prenames WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prename not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /return_ways
func UpdatePrename(c *gin.Context) {
	var prenames entity.Prename
	if err := c.ShouldBindJSON(&prenames); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", prenames.ID).First(&prenames); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prename not found"})
		return
	}

	if err := entity.DB().Save(&prenames).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prenames})
}
