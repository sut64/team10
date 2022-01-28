package controller

import (
	"net/http"


	"github.com/gin-gonic/gin"
	"github.com/sut64/team10/entity"
)

// POST /ProvinceController
func CreateProvince(c *gin.Context) {
	var provinces entity.Province
	if err := c.ShouldBindJSON(&provinces); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&provinces).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": provinces})
}

// GET /province/:id
func GetProvince(c *gin.Context) {
	var provinces entity.Province
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM provinces WHERE id = ?", id).Find(&provinces).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": provinces})
}

// GET /ProvinceController
func ListProvinces(c *gin.Context) {
	var provinces []entity.Province
	if err := entity.DB().Raw("SELECT * FROM provinces").Find(&provinces).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": provinces})
}

// DELETE /ProvinceController/:id
func DeleteProvince(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM provinces WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /ProvinceController
func UpdateProvince(c *gin.Context) {
	var province entity.Province
	if err := c.ShouldBindJSON(&province); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", province.ID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	if err := entity.DB().Save(&province).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": province})
}
