package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team10/entity"
)

func CreateTreatment(c *gin.Context) {

	var treatment entity.Treatmentrecord
	var disease entity.Disease
	var medicine entity.Medicine
	var personnel entity.Personnel
	var patientrecord entity.Patientrecord

	if err := c.ShouldBindJSON(&treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if tx := entity.DB().Where("id = ? ", treatment.PatientrecordID).First(&patientrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบผู้ป่วย"})
		return
	}

	if tx := entity.DB().Where("id = ? ", treatment.DiseaseID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาระบุโรค"})
		return
	}

	if tx := entity.DB().Where("id = ? ", treatment.MedicineID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "กรุณาระบุยา"})
		return
	}

	if tx := entity.DB().Where("id = ? ", treatment.PersonnelID).First(&personnel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบผู้ตรวจ"})
		return
	}

	tm := entity.Treatmentrecord{
		Disease:       disease,       // โยงความสัมพันธ์กับ Entity Disease
		Medicine:      medicine,      // โยงความสัมพันธ์กับ Entity Medicine
		Personnel:     personnel,     // โยงความสัมพันธ์กับ Entity Personnel
		Patientrecord: patientrecord, // โยงความสัมพันธ์กับ Entity PatientRecord

		// field Treatment
		Treatment:   treatment.Treatment,
		Temperature: treatment.Temperature,
		RecordDate:        treatment.RecordDate,
	}
	
	// validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(tm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save
	if err := entity.DB().Create(&tm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tm})

}

func GetTreatment(c *gin.Context) {

	var treatment entity.Treatmentrecord
	id := c.Param("id")
	if err := entity.DB().Preload("Disease").Preload("Medicine").Preload("Personnel").Preload("Patientrecord").Raw("Select * FROM treatmentrecords WHERE id = ? ", id).Scan(&treatment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})

}

func ListTreatment(c *gin.Context) {

	var treatment []entity.Treatmentrecord

	if err := entity.DB().Preload("Disease").Preload("Medicine").Preload("Personnel").Preload("Patientrecord").Raw("Select * FROM treatmentrecords ").Find(&treatment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})

}

func DeleteTreatment(c *gin.Context) {

	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM treatmentrecords WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatmentrecords not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

func UpdateTreatment(c *gin.Context) {

	var treatment entity.Treatmentrecord

	if err := c.ShouldBindJSON(&treatment); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", treatment.ID).First(&treatment); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "treatmentrecords not found"})

		return

	}

	if err := entity.DB().Save(&treatment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})

}
