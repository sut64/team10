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
		&Treatmentrecord{},
		&Appointment{})
	db = database

	disease1 := Disease{
		Diname: "ไม่มี",
	}
	db.Model(&Disease{}).Create(&disease1)

	disease2 := Disease{
		Diname: "ไข้หวัดใหญ่",
	}
	db.Model(&Disease{}).Create(&disease2)

	disease3 := Disease{
		Diname: "ไข้เลือดออก",
	}
	db.Model(&Disease{}).Create(&disease3)

	disease4 := Disease{
		Diname: "โรคเบาหวาน",
	}
	db.Model(&Disease{}).Create(&disease4)

	disease5 := Disease{
		Diname: "โรคภูมิแพ้",
	}
	db.Model(&Disease{}).Create(&disease5)

	medicine1 := Medicine{
		Medname:     "ยาแก้ไข้",
	}
	db.Model(&Medicine{}).Create(&medicine1)

	

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

	patientrecord1 := Patientrecord{
		Prename:        prename1,
		Firstname:      "กินอะไรดี",
		Lastname:       "อร่อยจัง",
		Gender:         gender1,
		Idcardnumber:   "1478899111112",
		Age:            20,
		Birthday:       time.Now(),
		BloodType:      bloodtype2,
		Phonenumber:    "0816629081",
		Email:          "eat@test.com",
		Home:           "111 moo2",
		Province:       province2,
		Emergencyname:  "มาสาย ลาก่อน",
		Emergencyphone: "0111111111",
		Timestamp:      time.Now(),
		Personnel:      personnel1,
	}
	db.Model(&Patientrecord{}).Create(&patientrecord1)


	db.Model(&Treatmentrecord{}).Create(&Treatmentrecord{
		Patientrecord: patientrecord1,
		Disease: disease2,
		Medicine: medicine1,
		Treatment: "การรักษาจึงเป็นเพียงการรักษาไปตามอาการเป็นสำคัญ กล่าวคือ ให้ยาลดไข้ เช็ดตัว ให้ดื่มน้ำมาก ๆ เพื่อป้องกันภาวะช็อก",
		Temperature: 32,
		Personnel:   personnel1,
		RecordDate:  time.Date(2022, 12, 22, 0, 0, 0, 0, time.UTC),
	})

	// DrugAllergy Data
	drugallergy1 := DrugAllergy{
		Name:    "แอสไพริน",
		Symptom: "หายใจลำบาก มีผื่นคัน",
	}
	db.Model(&DrugAllergy{}).Create(&drugallergy1)
	drugallergy2 := DrugAllergy{
		Name:    "เซรุ่มแก้พิษงู",
		Symptom: "เกิดผื่นลมพิษ กล้ามเนื้ออ่อนแรง คลื่นไส้ ปวดหัว ตาพร่ามัว หายใจลำบาก",
	}
	db.Model(&DrugAllergy{}).Create(&drugallergy2)
	drugallergy3 := DrugAllergy{
		Name:    "ยาชา",
		Symptom: "มีรอยช้ำ เลือดออก หรือเจ็บบริเวณที่ถูกฉีด",
	}
	db.Model(&DrugAllergy{}).Create(&drugallergy3)
}
