package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team10/entity"
)

// POST /drugallergys
func CreateDrugAllergy(c *gin.Context) {
	var drugallergy entity.DrugAllergy
	if err := c.ShouldBindJSON(&drugallergy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&drugallergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": drugallergy})
}

// GET /drugallergys/:id
func GetDrugAllergy(c *gin.Context) {
	var drugallergy entity.DrugAllergy
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM drug_allergies WHERE id = ?", id).Scan(&drugallergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": drugallergy})
}

// GET /drugallergys
func ListDrugAllergys(c *gin.Context) {
	var drugallergys []entity.DrugAllergy
	if err := entity.DB().Raw("SELECT * FROM drug_allergies").Scan(&drugallergys).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": drugallergys})
}

// DELETE /drugallergys/:id
func DeleteDrugAllergy(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM drug_allergies WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drug allergy not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /drugallergys
func UpdateDrugAllergy(c *gin.Context) {
	var drugallergy entity.DrugAllergy
	if err := c.ShouldBindJSON(&drugallergy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", drugallergy.ID).First(&drugallergy); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drug allergy not found"})
		return
	}
	if err := entity.DB().Save(&drugallergy).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": drugallergy})
}
