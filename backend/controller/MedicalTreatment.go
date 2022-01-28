package controller

import (
	"net/http"
	"github.com/sut64/team10/entity"


	"github.com/gin-gonic/gin"
)

// POST /MedicalTreatment
func CreateMedicalTreatment(c *gin.Context) {
	var treatment entity.MedicalTreatment
	if err := c.ShouldBindJSON(&treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": treatment})
}

// GET /MedicalTreatment/:id
func GetMedicalTreatment(c *gin.Context) {
	var treatment entity.MedicalTreatment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM medical_treatments WHERE id = ?", id).Scan(&treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})
}

// GET /MedicalTreatment
func ListMedicalTreatment(c *gin.Context) {
	var treatment []entity.MedicalTreatment
	if err := entity.DB().Raw("SELECT * FROM medical_treatments").Scan(&treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})
}

// DELETE /MedicalTreatment/:id
func DeleteMedicalTreatment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM medical_treatments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medical_treatments not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /MedicalTreatment
func UpdateMedicalTreatment(c *gin.Context) {
	var treatment entity.MedicalTreatment
	if err := c.ShouldBindJSON(&treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", treatment.ID).First(&treatment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medical_treatments not found"})
		return
	}

	if err := entity.DB().Save(&treatment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})
}
