
package controller

import (
	"net/http"
	"github.com/sut64/team10/entity"

	"github.com/gin-gonic/gin"
	"github.com/asaskevich/govalidator"
)

// POST /appointment
func CreateAppointment(c *gin.Context) {
	var appoint entity.Appointment
	var patient entity.Patientrecord
	var person entity.Personnel
	var treatment entity.Treatmentrecord

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร appoint
	if err := c.ShouldBindJSON(&appoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา Personnel ด้วย id
	if tx := entity.DB().Where("id = ?", appoint.PersonnelID).First(&person); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Personnel not found"})
		return
	}

	// 9: ค้นหา Patientrecord ด้วย id
	if tx := entity.DB().Where("id = ?", appoint.PatientrecordID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PatientRecord not found"})
		return
	}
	// 10: ค้นหา TreatmentRecord ด้วย id
	if tx := entity.DB().Where("id = ?", appoint.TreatmentrecordID).First(&treatment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TreatmentRecord not found"})
		return
	}

	// 12: สร้าง Appointment
	appointment := entity.Appointment{
		Personnel:       person,    // โยงความสัมพันธ์กับ Entity Personnel
		Patientrecord:   patient,   // โยงความสัมพันธ์กับ Entity PatientRecord
		Treatmentrecord: treatment, // โยงความสัมพันธ์กับ Entity TreatmentRecord
		Appoint_ID:      appoint.Appoint_ID,
		Room_number:     appoint.Room_number,
		Date_appoint:    appoint.Date_appoint,
	}

	//ขั้นตอน validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	// 13: บันทึก
	if err := entity.DB().Create(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": appointment})
}

// GET /appointment/:id
func GetAppointment(c *gin.Context) {
	var appoint entity.Appointment
	id := c.Param("id")
	if err := entity.DB().Preload("Personnel").Preload("Patientrecord").Preload("Treatmentrecord").Raw("SELECT * FROM appointments WHERE id = ?", id).Find(&appoint).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": appoint})
}

// GET /appointments
func ListAppointments(c *gin.Context) {
	var appointments []entity.Appointment
	if err := entity.DB().Preload("Personnel").Preload("Patientrecord").Preload("Treatmentrecord").Raw("SELECT * FROM appointments").Find(&appointments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": appointments})
}

// DELETE /appointment/:id
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM appointments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "appointment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /appointments
func UpdateAppointment(c *gin.Context) {
	var appointment entity.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", appointment.ID).First(&appointment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "appointment not found"})
		return
	}

	if err := entity.DB().Save(&appointment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": appointment})
}
