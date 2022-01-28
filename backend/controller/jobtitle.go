package controller

import (
	"net/http"
	"github.com/sut64/team10/entity"

	"github.com/gin-gonic/gin"
)

func ListJobTitle(c *gin.Context) {
	var jobtitle []entity.JobTitle
	if err := entity.DB().Raw("SELECT * FROM job_titles").Scan(&jobtitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobtitle})
}

func GetJobTitle(c *gin.Context) {
	var jobtitle entity.JobTitle
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM job_titles WHERE id = ?", id).Scan(&jobtitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobtitle})
}

func CreateJobTitle(c *gin.Context) {
	var jobtitle entity.JobTitle
	if err := c.ShouldBindJSON(&jobtitle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&jobtitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobtitle})
}

func UpdateJobTitle(c *gin.Context) {
	var jobtitle entity.JobTitle
	if err := c.ShouldBindJSON(&jobtitle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", jobtitle.ID).First(&jobtitle); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bloodtype not found"})
		return
	}

	if err := entity.DB().Save(&jobtitle).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jobtitle})
}

func DeleteJobTitle(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM job_titles WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "jobtitle not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
