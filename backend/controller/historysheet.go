package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team10/entity"
)

// POST /history_sheets
func CreateHistorySheet(c *gin.Context) {

	var historysheet entity.HistorySheet
	var patientrecord entity.Patientrecord
	var drugallergy entity.DrugAllergy
	var personnel entity.Personnel

	// ผลลัพธ์ที่ได้ ะถูก bind เข้าตัวแปร historysheet
	if err := c.ShouldBindJSON(&historysheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา personnel ด้วย id
	if tx := entity.DB().Where("id = ?", historysheet.PersonnelID).First(&personnel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "personnel not found"})
		return
	}

	// ค้นหา patientrecord ด้วย id
	if tx := entity.DB().Where("id = ?", historysheet.Patientrecord).First(&patientrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patientrecord not found"})
		return
	}

	// ค้นหา drugallergy ด้วย id
	if tx := entity.DB().Where("id = ?", historysheet.DrugAllergyID).First(&drugallergy); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "drug allergy not found"})
		return
	}
	// สร้าง HistorySheet
	hs := entity.HistorySheet{
		Personnel:     personnel,                // โยงความสัมพันธ์กับ Entity Personnel
		Patientrecord: patientrecord,            // โยงความสัมพันธ์กับ Entity Patientrecord
		DrugAllergy:   drugallergy,              // โยงความสัมพันธ์กับ Entity DrugAllergy
		Weight:        historysheet.Weight,      // ตั้งค่าฟิลด์ Weight
		Height:        historysheet.Height,      // ตั้งค่าฟิลด์ Height
		PressureOn:    historysheet.PressureOn,  // ตั้งค่าฟิลด์ PressureOn
		PressureLow:   historysheet.PressureLow, // ตั้งค่าฟิลด์ PressureLow
		Symptom:       historysheet.Symptom,     // ตั้งค่าฟิลด์ Symptom
	}

	// บันทึก
	if err := entity.DB().Create(&hs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hs})
}

// GET /historysheet/:id
func GetHistorySheet(c *gin.Context) {
	var historysheet entity.HistorySheet
	id := c.Param("id")
	if err := entity.DB().Preload("Personnel").Preload("Patientrecord").Preload("DrugAllergy").Raw("SELECT * FROM history_sheet WHERE id = ?", id).Find(&historysheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": historysheet})
}

// GET /history_sheets
func ListHistorySheets(c *gin.Context) {
	var historysheets []entity.HistorySheet
	if err := entity.DB().Preload("Personnel").Preload("Patientrecord").Preload("DrugAllergy").Raw("SELECT * FROM history_sheets").Find(&historysheets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": historysheets})
}

// DELETE /history_sheets/:id
func DeleteHistorySheet(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM history_sheets WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "history sheet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /history_sheets
func UpdateHistorySheet(c *gin.Context) {
	var historysheet entity.HistorySheet
	if err := c.ShouldBindJSON(&historysheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", historysheet.ID).First(&historysheet); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "historysheet not found"})
		return
	}

	if err := entity.DB().Save(&historysheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": historysheet})
}
