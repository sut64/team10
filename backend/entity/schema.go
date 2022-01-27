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
