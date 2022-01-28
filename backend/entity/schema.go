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

type MedicalTreatment struct {
	gorm.Model
	Tname string
	Price float64

}

type DrugAllergy struct {
	gorm.Model
	Name    string
	Symptom string
}

type HistorySheet struct {
	gorm.Model
	Weight      float64
	Height      float64
	Temperature float64
	PressureOn  uint
	PressureLow uint
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

type Personnel struct {
	gorm.Model
	Name        string
	Personalid  string
	BirthDay    time.Time
	Tel         string
	Address     string
	Salary      int
	GenderID    *uint
	Gender      Gender `gorm:"references:id"`
	BloodTypeID *uint
	BloodType   BloodType `gorm:"references:id"`
	JobTitleID  *uint
	JobTitle    JobTitle `gorm:"references:id"`
}
