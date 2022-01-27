package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Age       uint8
	BirthDay  time.Time
}

type Prename struct {
	gorm.Model
	Prename string
}

type Province struct {
	gorm.Model
	Province string
}

type Medicine struct {
	gorm.Model
	Medname     string
	Description string
	Quantity    string
	Price       float64
}
type DrugAllergy struct {
	gorm.Model
	Name    string
	Symptom string
}

type Gender struct {
	gorm.Model
	Genders string `gorm:"uniqueIndex"`
	//1 Gender เป็นเจ้าของได้หลาย Personnel_ID
	Personnels []Personnel `gorm:"foreignKey:GenderID"`
}

type BloodType struct {
	gorm.Model
	BloodType string `gorm:"uniqueIndex"`
	//1 BloodType เป็นเจ้าของได้หลาย Personnel_ID
	Personnels []Personnel `gorm:"foreignKey:BloodTypeID"`
}

type JobTitle struct {
	gorm.Model
	Job string `gorm:"uniqueIndex"`
	//1 JobTitle เป็นเจ้าของได้หลาย Personnel_ID
	Personnels []Personnel `gorm:"foreignKey:JobTitleID"`
}
