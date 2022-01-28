package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("se-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&User{},
		&Prename{},
		&Province{},
		&Medicine{},
		&MedicalTreatment{},
		&DrugAllergy{},
		&HistorySheet{},
		&Gender{},
		&BloodType{},
		&JobTitle{},
		&Personnel{},
		&Patientrecord{},
		&Bill{},
		&Disease{},
		&TreatmentRecord{},
		&Appointment{})
	db = database

	db.Model(&Medicine{}).Create(&Medicine{
		Medname:     "ไม่มี",
		Description: "none",
		Quantity:    "none",
		Price:       0,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Medname:     "Omepazole",
		Description: "none",
		Quantity:    "good",
		Price:       275.50,
	})
	db.Model(&Medicine{}).Create(&Medicine{
		Medname:     "ibupofen",
		Description: "none",
		Quantity:    "good",
		Price:       125.45,
	})

	db.Model(&MedicalTreatment{}).Create(&MedicalTreatment{
		Tname: "ไม่มี",
		Price: 0,
	})
	db.Model(&MedicalTreatment{}).Create(&MedicalTreatment{
		Tname: "Meet Doctor",
		Price: 250,
	})
	db.Model(&MedicalTreatment{}).Create(&MedicalTreatment{
		Tname: "X-ray",
		Price: 500,
	})
}
