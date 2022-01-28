package controller

import (
	"net/http"

	"github.com/VPNine/team10/entity"
	"github.com/gin-gonic/gin"
)

func CreateBill(c *gin.Context) {
	var bill entity.Bill
	var record entity.Patientrecord

	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bill.PatientrecordID).First(&record); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	// 10: สร้าง Bill
	bi := entity.Bill{
		Patientrecord: record,
		Cot:           bill.Cot,
		Com:           bill.Com,
		Sum:           bill.Com + bill.Cot,
		Listofbill:    bill.Listofbill,
		Dateofbill:    bill.Dateofbill,
	}

	// 11: บันทึก
	if err := entity.DB().Create(&bi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bi})
}

// GET /bill
// List all bill
func ListBill(c *gin.Context) {
	var bill []entity.Bill
	if err := entity.DB().Preload("Patientrecord").Raw("SELECT * FROM bills").Find(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// GET /bill/:id
// Get bill by id

func GetBill(c *gin.Context) {
	var bill entity.Bill
	id := c.Param("id")
	if err := entity.DB().Preload("Patientrecord").Preload("MedicalTreatment").Preload("Medicine").Raw("SELECT * FROM bills WHERE id = ?", id).Scan(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// PATCH /Bill
func UpdateBill(c *gin.Context) {
	var bill entity.Bill
	if err := c.ShouldBindJSON(&bill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bill.ID).First(&bill); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}

	if err := entity.DB().Save(&bill).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bill})
}

// DELETE /Bill/:id
func DeleteBill(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM bills WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bill not found"})
		return
	}
	/*
		if err := entity.DB().Where("id = ?", id).Delete(&entity.billinfo{}).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}*/

	c.JSON(http.StatusOK, gin.H{"data": id})
}
