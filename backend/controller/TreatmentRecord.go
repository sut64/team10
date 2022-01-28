package controller

import (
	"net/http"

	"github.com/Thanawat-Launakorn/treatment/entity"
	"github.com/gin-gonic/gin"
)

func CreateTreatment(c *gin.Context) {

	var treatment entity.TreatmentRecord
	var disease entity.Disease
	var medicine entity.Medicine
	var personnel entity.Personnel
	var patientrecord entity.PatientRecord

	if err := c.ShouldBindJSON(&treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ? ", treatment.DiseaseID).First(&disease); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "disease not found"})
		return
	}

	if tx := entity.DB().Where("id = ? ", treatment.MedicineID).First(&medicine); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "medicine not found"})
		return
	}

	if tx := entity.DB().Where("id = ? ", treatment.PatientRecordID).First(&patientrecord); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patienrecord not found"})
		return
	}

	if tx := entity.DB().Where("id = ? ", treatment.PersonnelID).First(&personnel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "personnel not found"})
		return
	}

	tm := entity.TreatmentRecord{
		Disease:       disease,       // โยงความสัมพันธ์กับ Entity Disease
		Medicine:      medicine,      // โยงความสัมพันธ์กับ Entity Medicine
		Personnel:     personnel,     // โยงความสัมพันธ์กับ Entity Personnel
		PatientRecord: patientrecord, // โยงความสัมพันธ์กับ Entity PatientRecord

		// field Treatment
		Treatment:   treatment.Treatment,
		Temperature: treatment.Temperature,
		Date:        treatment.Date,
	}
	if err := entity.DB().Create(&tm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": tm})

}

func GetTreatment(c *gin.Context) {

	var treatment entity.TreatmentRecord
	id := c.Param("id")
	if err := entity.DB().Preload("Disease").Preload("Medicine").Preload("Personnel").Preload("PatientRecord").Raw("Select * FROM treatment_records WHERE id = ? ", id).Scan(&treatment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})

}

func ListTreatment(c *gin.Context) {

	var treatment []entity.TreatmentRecord

	if err := entity.DB().Preload("Disease").Preload("Medicine").Preload("Personnel").Preload("PatientRecord").Raw("Select * FROM treatment_records ").Find(&treatment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})

}

func DeleteTreatment(c *gin.Context) {

	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM treatment_records WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "treatment_records not found"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": id})

}

func UpdateTreatment(c *gin.Context) {

	var treatment entity.TreatmentRecord

	if err := c.ShouldBindJSON(&treatment); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", treatment.ID).First(&treatment); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "treatment_records not found"})

		return

	}

	if err := entity.DB().Save(&treatment).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": treatment})

}

