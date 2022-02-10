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
	Weight      float32 `valid:"positiveFloat~Weight must be positive"`
	Height      float32
	Temperature float32
	PressureOn  int `valid:"positiveIntOn~Pressure on must be positive"`
	PressureLow uint
	Symptom     string `valid:"required~Symptom not blank"`

	//patientrecord_id ทำหน้าที่เป็น FK
	PatientrecordID *uint
	Patientrecord   Patientrecord `gorm:"references:id" valid:"-"`

	//Personnel_id ทำหน้าที่เป็น FK
	PersonnelID *uint
	Personnel   Personnel `gorm:"references:id" valid:"-"`

	//DrugAllergy_id ทำหน้าที่เป็น FK
	DrugAllergyID *uint
	DrugAllergy   DrugAllergy `gorm:"references:id" valid:"-"`
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
	Personalid string    `valid:"matches(^[0123456789]{13}$)~กรุณากรอกบัตรประจำตัวประชาชนให้ถูกต้อง"`
	BirthDay   time.Time `valid:"past~กรุณากรอกวันเกิดให้ถูกต้อง"`
	Tel        string
	Address    string
	Salary     int `valid:"positiveIntSalaryForPersonnel~กรุณาเงินเดือนให้มีค่าเป็นบวก"`

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

	Firstname string
	//`valid:"required~Firstname cannot be blank."`
	Lastname string
	//`valid:"required~Lastname cannot be blank."`
	Idcardnumber string    `valid:"matches(^[1-9]\\d{12}$)~Idcardnumber must be 0 - 19 and contain 13 digits."`
	Age          int       `valid:"intnotlessthanZero~Age must not be less than 0."`
	Birthday     time.Time `valid:"past~Birthday must be in the past."`
	Phonenumber  string
	//`valid:"required~Phonenumber cannot be blank."`
	Email string
	Home  string
	//`valid:"required~Address cannot be blank."`
	Emergencyname  string
	Emergencyphone string
	Timestamp      time.Time

	//prename_id ทำหน้าที่เป็น FK
	PrenameID *uint
	Prename   Prename `gorm:"references:id" valid:"-"`

	//gender_id ทำหน้าที่เป็น FK
	GenderID *uint
	Gender   Gender `gorm:"references:id" valid:"-"`

	//bloodtype_id ทำหน้าที่เป็น FK
	BloodTypeID *uint
	BloodType   BloodType `gorm:"references:id" valid:"-"`

	//province_id ทำหน้าที่เป็น FK
	ProvinceID *uint
	Province   Province `gorm:"references:id" valid:"-"`

	//personnel_id ทำหน้าที่เป็น FK
	PersonnelID *uint
	Personnel   Personnel `gorm:"references:id" valid:"-"`

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
	Listofbill int       `valid:"intnotZero,required~List must not Zero"`
	Dateofbill time.Time `valid:"future~ไม่สามารถบันทึกเป็นเวลาในอดีตได้"`

	PatientrecordID *uint
	Patientrecord   Patientrecord `gorm:"references:id" valid:"-"`

	MedicineID *uint
	Medicine   Medicine `gorm:"references:id" valid:"-"`

	MedicalTreatmentID *uint
	MedicalTreatment   MedicalTreatment `gorm:"references:id" valid:"-"`
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
	Temperature int       `valid:"positiveInt~อุณหภูมิควรจะอยู่ในช่วงของ 32 - 40, required~อุณหภูมิควรจะอยู่ในช่วงของ 32 - 40" `
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
	Appoint_ID   string    `valid:"required,matches(^[A]\\d{4}$)~Appoint ID must Start with A and (0-9) 4 digits"`
	Room_number  int       `valid:"required,positiveIntRoomNumber~Room_number: non zero value required"`
	Date_appoint time.Time `valid:"required,future~Date Appointment must be in the future"`

	PatientrecordID *uint
	Patientrecord   Patientrecord `gorm:"references:id" valid:"-"`

	PersonnelID *uint
	Personnel   Personnel `gorm:"references:id" valid:"-"`

	TreatmentRecordID *uint
	TreatmentRecord   TreatmentRecord `gorm:"references:id" valid:"-"`
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
	govalidator.CustomTypeTagMap.Set("positiveIntSalaryForPersonnel", func(i interface{}, context interface{}) bool {
		v, _ := i.(int)
		return govalidator.InRangeInt(v, 1, 1000000)
	})

	govalidator.CustomTypeTagMap.Set("positiveIntRoomNumber", func(i interface{}, context interface{}) bool {
		v, _ := i.(int)
		return govalidator.InRangeInt(v, 1, 10)
	})

	govalidator.CustomTypeTagMap.Set("intnotlessthanZero", func(i interface{}, context interface{}) bool {
		v, _ := i.(int)
		return govalidator.InRangeInt(v, 0, 999999)
	})
	govalidator.CustomTypeTagMap.Set("positiveIntOn", func(i interface{}, context interface{}) bool {
		v, _ := i.(int)
		return govalidator.InRangeInt(v, 1, 9999)
	})
}
