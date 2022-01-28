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
	//1 Prename เป็นเจ้าของได้หลาย patientrecord_ID
	Patientrecord []Patientrecord `gorm:"foreignKey:PrenameID"`
}

type Province struct {
	gorm.Model
	Province string
	//1 Province เป็นเจ้าของได้หลาย patientrecord_ID
	Patientrecord []Patientrecord `gorm:"foreignKey:ProvinceID"`
}

type Medicine struct {
	gorm.Model
	Medname     string
	Description string
	Quantity    string
	Price       float64

	Bills []Bill `gorm:"foreignKey:MedicineID"`
}

type MedicalTreatment struct {
	gorm.Model
	Tname string
	Price float64

	Bills []Bill `gorm:"foreignKey:MedicalTreatmentID"`
}

type DrugAllergy struct {
	gorm.Model
	Name    string
	Symptom string

	//1 DrugAllergy เป็นเจ้าของได้หลาย HistorySheet_ID
	HistorySheet []HistorySheet `gorm:"foreignKey:DrugAllergyID"`
}

type HistorySheet struct {
	gorm.Model
	Weight      float64
	Height      float64
	Temperature float64
	PressureOn  uint
	PressureLow uint
	Symptom     string

	//patientrecord_id ทำหน้าที่เป็น FK
	PatientrecordID *uint
	Patientrecord   Patientrecord `gorm:"references:id"`

	//Personnel_id ทำหน้าที่เป็น FK
	PersonnelID *uint
	Personnel   Personnel `gorm:"references:id"`

	//DrugAllergy_id ทำหน้าที่เป็น FK
	DrugAllergyID *uint
	DrugAllergy   DrugAllergy `gorm:"references:id"`
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

	//1 Personnel เป็นเจ้าของได้หลาย HistorySheet_ID
	HistorySheet []HistorySheet `gorm:"foreignKey:PersonnelID"`

	Appointments []Appointment `gorm:"foreignKey:PersonnelID"`
}

type Patientrecord struct {
	gorm.Model

	Firstname      string
	Lastname       string
	Idcardnumber   string
	Age            uint8
	Birthday       time.Time
	Phonenumber    string
	Email          string
	Home           string
	Emergencyname  string
	Emergencyphone string
	Timestamp      time.Time

	//prename_id ทำหน้าที่เป็น FK
	PrenameID *uint
	Prename   Prename `gorm:"references:id"`

	//gender_id ทำหน้าที่เป็น FK
	GenderID *uint
	Gender   Gender `gorm:"references:id"`

	//bloodtype_id ทำหน้าที่เป็น FK
	BloodtypeID *uint
	Bloodtype   BloodType `gorm:"references:id"`

	//province_id ทำหน้าที่เป็น FK
	ProvinceID *uint
	Province   Province `gorm:"references:id"`

	//personnel_id ทำหน้าที่เป็น FK
	PersonnelID *uint
	Personnel   Personnel `gorm:"references:id"`

	Bills []Bill `gorm:"foreignKey:PatientrecordID"`

	//1 Patientrecord เป็นเจ้าของได้หลาย HistorySheet_ID
	HistorySheet []HistorySheet `gorm:"foreignKey:PatientrecordID"`

	Appointments []Appointment `gorm:"foreignKey:PatientrecordID"`
}

type Bill struct {
	gorm.Model
	Cot        float64   `valid:"-"`
	Com        float64   `valid:"-"`
	Sum        float64   `valid:"positiveFloat,required~Sum must not Zero"`
	Listofbill int       `valid:"positiveInt,required~List must not Zero"`
	Dateofbill time.Time `valid:"future"`

	PatientrecordID *uint
	Patientrecord   Patientrecord `gorm:"references:id" valid:"-"`

	MedicineID *uint
	Medicine   Medicine `gorm:"references:id" valid:"-"`

	MedicalTreatmentID *uint
	MedicalTreatment   MedicalTreatment `gorm:"references:id" valid:"-"`

}

type Disease struct {

	gorm.Model
	Diname string
	Description string
	
	//1 Di เป็นเจ้าของได้หลาย TreatmentRecord_ID
	TreatmentRecords []TreatmentRecord `gorm"foreignKey:DiseaseID"`
}


}

type TreatmentRecord struct {

	gorm.Model
	Treatment string 
	Temperature float32
	Date time.Time	
	
	PersonnelID *uint
	Personnel Personnel `gorm:"reference:id"`
	
	PatientrecordID *uint
	Patientrecord Patientrecord `gorm:"reference:id"`

	MedicineID *uint
	Medicine Medicine `gorm:"reference:id"`

	DiseaseID *uint
	Disease Disease  `gorm:"reference:id"`

	Appointments []Appointment `gorm:"foreignKey:TreatmentRecordID"`
}

type Appointment struct {
	gorm.Model
	Appoint_ID   string
	Room_number  uint
	Date_appoint time.Time

	PatientrecordID *uint
	Patientrecord   Patientrecord `gorm:"references:id"`

	PersonnelID *uint
	Personnel   Personnel `gorm:"references:id"`

	TreatmentRecordID *uint
	TreatmentRecord   TreatmentRecord `gorm:"references:id"`

}

type Medicine struct {

	gorm.Model
	MedName string
	Description string

	//1 Med เป็นเจ้าของได้หลาย TreatmentRecord_ID
	TreatmentRecords []TreatmentRecord  `gorm"foreignKey:MedicineID"`

}




