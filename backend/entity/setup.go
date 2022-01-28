package entity

import (
	"time"

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

	db.Model(&Gender{}).Create(&Gender{
		Genders: "ชาย",
	})
	db.Model(&Gender{}).Create(&Gender{
		Genders: "หญิง",
	})

	db.Model(&BloodType{}).Create(&BloodType{
		BloodType: "A",
	})
	db.Model(&BloodType{}).Create(&BloodType{
		BloodType: "B",
	})
	db.Model(&BloodType{}).Create(&BloodType{
		BloodType: "AB",
	})
	db.Model(&BloodType{}).Create(&BloodType{
		BloodType: "O",
	})

	db.Model(&JobTitle{}).Create(&JobTitle{
		Job: "หมอ",
	})
	db.Model(&JobTitle{}).Create(&JobTitle{
		Job: "พยาบาล",
	})
	db.Model(&JobTitle{}).Create(&JobTitle{
		Job: "ผู้ช่วยพยาบาล",
	})
	db.Model(&JobTitle{}).Create(&JobTitle{
		Job: "ทันตแพทย์",
	})
	db.Model(&JobTitle{}).Create(&JobTitle{
		Job: "เภสัชกร",
	})
	db.Model(&JobTitle{}).Create(&JobTitle{
		Job: "เทคนิคการแพทย์",
	})
	db.Model(&JobTitle{}).Create(&JobTitle{
		Job: "พนักงานบัญชี",
	})
	db.Model(&JobTitle{}).Create(&JobTitle{
		Job: "ภารโรง",
	})
	db.Model(&JobTitle{}).Create(&JobTitle{
		Job: "รปภ.",
	})

	// Prename Data
	prename1 := Prename{
		Prename: "นาย",
	}
	db.Model(&Prename{}).Create(&prename1)

	prename2 := Prename{
		Prename: "นาง",
	}
	db.Model(&Prename{}).Create(&prename2)

	prename3 := Prename{
		Prename: "นางสาว",
	}
	db.Model(&Prename{}).Create(&prename3)

	// Province Data
	province1 := Province{
		Province: "นครราชสีมา",
	}
	db.Model(&Province{}).Create(&province1)

	province2 := Province{
		Province: "อุบลราชธานี",
	}
	db.Model(&Province{}).Create(&province2)

	province3 := Province{
		Province: "กรุงเทพ",
	}
	db.Model(&Province{}).Create(&province3)

	// JobTitle Data
	job1 := JobTitle{
		Job: "พนักงานเวชระเบียน",
	}
	db.Model(&JobTitle{}).Create(&job1)

	gender1 := Gender{
		Genders: "ชาย",
	}
	db.Model(&Gender{}).Create(&gender1)

	bloodtype2 := BloodType{
		BloodType: "B",
	}
	db.Model(&BloodType{}).Create(&bloodtype2)

	// Personnel Data
	personnel1 := Personnel{
		Name:       "ขยัน อดทด",
		Personalid: "ABABAB",
		BirthDay:   time.Now(),
		Tel:        "0555555555555",
		Address:    "1111",
		Salary:     20,
		Gender:     gender1,
		BloodType:  bloodtype2,
		JobTitle:   job1,
	}
	db.Model(&Personnel{}).Create(&personnel1)

	// Patientrecord 1
	db.Model(&Patientrecord{}).Create(&Patientrecord{
		Prename:        prename1,
		Firstname:      "นคร",
		Lastname:       "ศรีสรรณ์",
		Gender:         gender1,
		Idcardnumber:   "1234455678948",
		Age:            25,
		Birthday:       time.Now(),
		BloodType:      bloodtype2,
		Phonenumber:    "0855555555",
		Email:          "nakorn@test.com",
		Home:           "111 moo1",
		Province:       province1,
		Emergencyname:  "มาสาย ลาก่อน",
		Emergencyphone: "0111111111",
		Timestamp:      time.Now(),
		Personnel:      personnel1,
	})

}
