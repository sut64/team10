package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
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

	TreatmentRecords []TreatmentRecord `gorm"foreignKey:MedicineID"`
	Bills            []Bill            `gorm:"foreignKey:MedicineID"`
}

type MedicalTreatment struct {
	gorm.Model
	Tname string
	Price float32

	Bills []Bill `gorm:"foreignKey:MedicalTreatmentID"`
}

type DrugAllergy struct {
	gorm.Model
	Name    string
	Symptom string

	//1 DrugAllergy เป็นเจ้าของได้หลาย HistorySheet_ID
	HistorySheets []HistorySheet `gorm:"foreignKey:DrugAllergyID"`
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

	Patientrecord []Patientrecord `gorm:"foreignKey:GenderID"`
}

type BloodType struct {
	gorm.Model
	BloodType string `gorm:"uniqueIndex"`
	//1 BloodType เป็นเจ้าของได้หลาย Personnel_ID
	Personnels []Personnel `gorm:"foreignKey:BloodTypeID"`

	Patientrecord []Patientrecord `gorm:"foreignKey:BloodTypeID"`
}

type JobTitle struct {
	gorm.Model
	Job string `gorm:"uniqueIndex"`
	//1 JobTitle เป็นเจ้าของได้หลาย Personnel_ID
	Personnels []Personnel `gorm:"foreignKey:JobTitleID"`
}

type Personnel struct {
	gorm.Model
	Name       string
	Personalid string
	BirthDay   time.Time
	Tel        string
	Address    string
	Salary     int

	GenderID *uint
	Gender   Gender `gorm:"references:id"`

	BloodTypeID *uint
	BloodType   BloodType `gorm:"references:id"`

	JobTitleID *uint
	JobTitle   JobTitle `gorm:"references:id"`

	//1 Personnel เป็นเจ้าของได้หลาย HistorySheet_ID
	HistorySheets    []HistorySheet    `gorm:"foreignKey:PersonnelID"`
	TreatmentRecords []TreatmentRecord `gorm"foreignKey:PersonnelID"`
	Appointments     []Appointment     `gorm:"foreignKey:PersonnelID"`
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
	BloodTypeID *uint
	BloodType   BloodType `gorm:"references:id"`

	//province_id ทำหน้าที่เป็น FK
	ProvinceID *uint
	Province   Province `gorm:"references:id"`

	//personnel_id ทำหน้าที่เป็น FK
	PersonnelID *uint
	Personnel   Personnel `gorm:"references:id"`

	//1 Patientrecord เป็นเจ้าของได้หลาย HistorySheet_ID
	HistorySheets    []HistorySheet    `gorm:"foreignKey:PatientrecordID"`
	TreatmentRecords []TreatmentRecord `gorm"foreignKey:PatientrecordID"`
	Appointments     []Appointment     `gorm:"foreignKey:PatientrecordID"`
	Bills            []Bill            `gorm:"foreignKey:PatientrecordID"`
}

type Bill struct {
	gorm.Model
	Cot        float32
	Com        float32
	Sum        float32   `valid:"positiveFloat,required~Sum must not Zero"`
	Listofbill int	`valid:"intnotZero,required~List must not Zero"`
	Dateofbill time.Time `valid:"future~ไม่สามารถบันทึกเป็นเวลาในอดีตได้"`

	PatientrecordID *uint
	Patientrecord   Patientrecord `gorm:"references:id"`

	MedicineID *uint
	Medicine   Medicine `gorm:"references:id"`

	MedicalTreatmentID *uint
	MedicalTreatment   MedicalTreatment `gorm:"references:id"`
}

type Disease struct {
	gorm.Model
	Diname      string
	Description string

	//1 Di เป็นเจ้าของได้หลาย TreatmentRecord_ID
	TreatmentRecords []TreatmentRecord `gorm"foreignKey:DiseaseID"`
}

type TreatmentRecord struct {
	gorm.Model
	Treatment   string    `valid:"required~กรุณากรอกวิธีการรักษา"`
	Temperature int       `valid:"positiveInt~อุณหภูมิควรจะอยู่ในช่วงของ 32 - 40" `
	RecordDate  time.Time `valid:"future~ไม่สามารถบันทึกเป็นเวลาในอดีตได้"`

	PersonnelID *uint
	Personnel   Personnel `gorm:"reference:id" valid:"-"`

	PatientrecordID *uint
	Patientrecord   Patientrecord `gorm:"reference:id" valid:"-"`

	MedicineID *uint
	Medicine   Medicine `gorm:"reference:id" valid:"-"`

	DiseaseID *uint
	Disease   Disease `gorm:"reference:id" valid:"-"`

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

func init() {
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now())
	})

	govalidator.CustomTypeTagMap.Set("positiveInt", func(i interface{}, context interface{}) bool {
		v, _ := i.(int)
		return govalidator.InRangeInt(v, 32, 40)
	})
	govalidator.CustomTypeTagMap.Set("intnotZero", func(i interface{}, context interface{}) bool {
		v, _ := i.(int)
		return govalidator.InRangeInt(v, 1, 999)
	})
	govalidator.CustomTypeTagMap.Set("positiveFloat", func(i interface{}, context interface{}) bool {
		v, _ := i.(float32)
		return govalidator.InRangeFloat32(v, 1.00, 9999.99)
	})
}
