package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/sut64/team10/entity"

	"github.com/gin-gonic/gin"
)

func CreatePersonnel(c *gin.Context) {
	var personnel entity.Personnel
	var bloodtype entity.BloodType
	var gender entity.Gender
	var jobtitle entity.JobTitle

	if err := c.ShouldBindJSON(&personnel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", personnel.BloodTypeID).First(&bloodtype); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", personnel.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Member not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", personnel.JobTitleID).First(&jobtitle); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Btype not found"})
		return
	}
	ps := entity.Personnel{
		BloodType:  bloodtype,
		Gender:     gender,
		JobTitle:   jobtitle,
		Name:       personnel.Name,
		Personalid: personnel.Personalid,
		BirthDay:   personnel.BirthDay,
		Tel:        personnel.Tel,
		Address:    personnel.Address,
		Salary:     personnel.Salary,
	}

	//ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(ps); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&ps).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ps})
}

func ListPersonnel(c *gin.Context) {
	var personnel []entity.Personnel
	if err := entity.DB().Preload("BloodType").Preload("Gender").Preload("JobTitle").Raw("SELECT * FROM personnels").Find(&personnel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": personnel})
}

func GetPersonnel(c *gin.Context) {
	var personnel entity.Personnel
	id := c.Param("id")
	if err := entity.DB().Preload("BloodType").Preload("Gender").Preload("JobTitle").Raw("SELECT * FROM personnels WHERE id = ?", id).Scan(&personnel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": personnel})
}

func UpdatePersonnel(c *gin.Context) {
	var personnel entity.Personnel
	if err := c.ShouldBindJSON(&personnel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", personnel.ID).First(&personnel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "personnel not found"})
		return
	}

	if err := entity.DB().Save(&personnel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": personnel})
}

func DeletePersonnel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM personnels WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "personnel not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
