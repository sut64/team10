package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPatientrecordPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	patientrecord := Patientrecord{
		Firstname:      "Nakorn",
		Lastname:       "Srisan",
		Idcardnumber:   "1234567891234",
		Age:            25,
		Birthday:       time.Date(2010, 1, 15, 0, 0, 0, 0, time.UTC),
		Phonenumber:    "0855555555",
		Email:          "nakorn@test.com",
		Home:           "111 moo1",
		Emergencyname:  "Masai Lagon",
		Emergencyphone: "0111111111",
		Timestamp:      time.Now(),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(patientrecord)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

// ตรวจสอบหมายเลข Idcardnumber ต้องขึ้นด้วย 1 ถึง 9 และตามด้วย 0 ถึง 9 จำนวน 12 ตัว
func TestIdcardnumberMustBeValid(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"123456789012",   // 12 ตัว
		"12345678901234", // 14 ตัว
		"0000000000000",  // ตัวแรกเป็น 0
		"10000000000x",
		"x0000000000x",
		"0",
		"xxxxxxxxxxxxx",
	}

	for _, fixture := range fixtures {
		patientrecord := Patientrecord{
			Firstname:      "Nakorn",
			Lastname:       "Srisan",
			Idcardnumber:   fixture, //ผิด
			Age:            25,
			Birthday:       time.Date(2010, 1, 15, 0, 0, 0, 0, time.UTC),
			Phonenumber:    "0855555555",
			Email:          "nakorn@test.com",
			Home:           "111 moo1",
			Emergencyname:  "Masai Lagon",
			Emergencyphone: "0111111111",
			Timestamp:      time.Now(),
		}

		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(patientrecord)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Idcardnumber must be 0 - 9 and contain 13 digits."))
	}
}

// ตรวจสอบอายุต้องเป็นตัวเลขที่ไม่น้อยกว่า 0
func TestAgenotlessthanZero(t *testing.T) {
	g := NewGomegaWithT(t)

	patientrecord := Patientrecord{
		Firstname:      "Nakorn",
		Lastname:       "Srisan",
		Idcardnumber:   "1234567891234",
		Age:            -1, //ผิด
		Birthday:       time.Date(2010, 1, 15, 0, 0, 0, 0, time.UTC),
		Phonenumber:    "0855555555",
		Email:          "nakorn@test.com",
		Home:           "111 moo1",
		Emergencyname:  "Masai Lagon",
		Emergencyphone: "0111111111",
		Timestamp:      time.Now(),
	}

	ok, err := govalidator.ValidateStruct(patientrecord)

	// ok ต้องไม่เป็น true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็น nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error() ต้องมี message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Age must not be less than 0."))

}

// ตรวจสอบวันเดือนปีเกิดต้องเป็นวันในอดีต
func TestBirthdateMustBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	patientrecord := Patientrecord{
		Firstname:      "Nakorn",
		Lastname:       "Srisan",
		Idcardnumber:   "1234567891234",
		Age:            25,
		Birthday:       time.Now().Add(24 * time.Hour), // อนาคต, fail
		Phonenumber:    "0855555555",
		Email:          "nakorn@test.com",
		Home:           "111 moo1",
		Emergencyname:  "Masai Lagon",
		Emergencyphone: "0111111111",
		Timestamp:      time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(patientrecord)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Birthday must be in the past."))

}

// ตรวจสอบ ชื่อ ต้องไม่เว้นว่าง
func TestFirstnameNotBlankt(t *testing.T) {
	g := NewGomegaWithT(t)

	patientrecord := Patientrecord{
		Firstname:      "", //ผิด
		Lastname:       "Srisan",
		Idcardnumber:   "1234567891234",
		Age:            25,
		Birthday:       time.Date(2010, 1, 15, 0, 0, 0, 0, time.UTC),
		Phonenumber:    "0855555555",
		Email:          "nakorn@test.com",
		Home:           "111 moo1",
		Emergencyname:  "Masai Lagon",
		Emergencyphone: "0111111111",
		Timestamp:      time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(patientrecord)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Firstname cannot be blank."))
}

// ตรวจสอบ นามสกุล ต้องไม่เว้นว่าง
func TestLastnameNotBlankt(t *testing.T) {
	g := NewGomegaWithT(t)

	patientrecord := Patientrecord{
		Firstname:      "Nakorn",
		Lastname:       "", //ผิด
		Idcardnumber:   "1234567891234",
		Age:            25,
		Birthday:       time.Date(2010, 1, 15, 0, 0, 0, 0, time.UTC),
		Phonenumber:    "0855555555",
		Email:          "nakorn@test.com",
		Home:           "111 moo1",
		Emergencyname:  "Masai Lagon",
		Emergencyphone: "0111111111",
		Timestamp:      time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(patientrecord)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Lastname cannot be blank."))
}

// ตรวจสอบ ที่อยู่ ต้องไม่เว้นว่าง
func TestHomeNotBlankt(t *testing.T) {
	g := NewGomegaWithT(t)

	patientrecord := Patientrecord{
		Firstname:      "Nakorn",
		Lastname:       "Srisan",
		Idcardnumber:   "1234567891234",
		Age:            25,
		Birthday:       time.Date(2010, 1, 15, 0, 0, 0, 0, time.UTC),
		Phonenumber:    "0855555555",
		Email:          "nakorn@test.com",
		Home:           "", //ผิด
		Emergencyname:  "Masai Lagon",
		Emergencyphone: "0111111111",
		Timestamp:      time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(patientrecord)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Address cannot be blank."))
}
