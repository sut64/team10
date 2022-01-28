package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team10/entity"
)

// POST /Patientrecords
func CreatePatientrecord(c *gin.Context) {

	var patientrecord entity.Patientrecord
	var bloodtype entity.Bloodtype
	var gender entity.Gender
	var personnel entity.Personnel
	var prename entity.Prename
	var province entity.Province

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร Patientrecord
	if err := c.ShouldBindJSON(&patientrecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา bloodtype ด้วย id
	if tx := entity.DB().Where("id = ?", patientrecord.BloodtypeID).First(&bloodtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bloodtype not found"})
		return
	}

	// 10: ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", patientrecord.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// 11: ค้นหา personnel ด้วย id
	if tx := entity.DB().Where("id = ?", patientrecord.PersonnelID).First(&personnel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "personnel not found"})
		return
	}

	// 12: ค้นหา prename ด้วย id
	if tx := entity.DB().Where("id = ?", patientrecord.PrenameID).First(&prename); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prename not found"})
		return
	}

	// 13: ค้นหา province ด้วย id
	if tx := entity.DB().Where("id = ?", patientrecord.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	// 12: สร้าง Patientrecord
	rb := entity.Patientrecord{
		Prename:        prename,                      // โยงความสัมพันธ์กับ Entity Prename
		Firstname:      patientrecord.Firstname,      // ตั้งค่าฟิลด์ Firstname
		Lastname:       patientrecord.Lastname,       // ตั้งค่าฟิลด์ Lastname
		Gender:         gender,                       // โยงความสัมพันธ์กับ Entity Gender
		Idcardnumber:   patientrecord.Idcardnumber,   // ตั้งค่าฟิลด์ Idcardnumber
		Age:            patientrecord.Age,            // ตั้งค่าฟิลด์ Age
		Birthday:       patientrecord.Birthday,       // ตั้งค่าฟิลด์ Birthday
		Bloodtype:      bloodtype,                    // โยงความสัมพันธ์กับ Entity BloodType
		Phonenumber:    patientrecord.Phonenumber,    // ตั้งค่าฟิลด์ Phonenumber
		Email:          patientrecord.Email,          // ตั้งค่าฟิลด์ Email
		Home:           patientrecord.Home,           // ตั้งค่าฟิลด์ Home
		Province:       province,                     // โยงความสัมพันธ์กับ Entity Province
		Emergencyname:  patientrecord.Emergencyname,  // ตั้งค่าฟิลด์ Emergencyname
		Emergencyphone: patientrecord.Emergencyphone, // ตั้งค่าฟิลด์ Emergencyphone
		Timestamp:      patientrecord.Timestamp,      // ตั้งค่าฟิลด์ Timestamp
		Personnel:      personnel,                    // โยงความสัมพันธ์กับ Entity Personnel
	}

	// 13: บันทึก
	if err := entity.DB().Create(&rb).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rb})
}

// GET /patientrecord/:id
func GetPatientrecord(c *gin.Context) {
	var patientrecords entity.Patientrecord
	id := c.Param("id")
	if err := entity.DB().Preload("BloodType").Preload("Gender").Preload("Personnel").Preload("Prename").Preload("Province").Raw("SELECT * FROM patientrecords WHERE id = ?", id).Find(&patientrecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": patientrecords})
}

// GET /patientrecords
func ListPatientrecords(c *gin.Context) {
	var patientrecords []entity.Patientrecord
	if err := entity.DB().Preload("BloodType").Preload("Gender").Preload("Personnel").Preload("Prename").Preload("Province").Raw("SELECT * FROM patientrecords").Find(&patientrecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patientrecords})
}

// DELETE /patientrecords/:id
func DeletePatientrecord(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM patientrecords WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patientrecord not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /patientrecords
func UpdatePatientrecord(c *gin.Context) {
	var patientrecords entity.Patientrecord
	if err := c.ShouldBindJSON(&patientrecords); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", patientrecords.ID).First(&patientrecords); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patientrecord not found"})
		return
	}

	if err := entity.DB().Save(&patientrecords).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": patientrecords})
}
