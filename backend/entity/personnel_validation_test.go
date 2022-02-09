package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPersonnelPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	personnel := Personnel{
		Name:       "Poramee Suriyajanno",
		Personalid: "1104200227093",
		BirthDay:   time.Date(2001, 8, 2, 0, 0, 0, 0, time.UTC),
		Tel:        "0638267373",
		Address:    "889 หมู่ 1 ต.ระแงง อ.ศีขรภูมิ จ.สุรินทร์ 32110",
		Salary:     15000,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(personnel)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestPersonalIDMustBeInValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	personnel := Personnel{
		Name:       "Poramee",
		Personalid: "11042002270934", // fail มีตัวอักษรที่นอกเหนือจาอก 0-9 อยู่ข้างใน
		BirthDay:   time.Date(2001, 8, 2, 0, 0, 0, 0, time.UTC),
		Tel:        "0638267373",
		Address:    "889 หมู่ 1 ต.ระแงง อ.ศีขรภูมิ จ.สุรินทร์ 32110",
		Salary:     15000,
	}

	ok, err := govalidator.ValidateStruct(personnel)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Personalid: does not validate as matches(^[0123456789]{13}$)"))
}

func TestBirthDayMustBePast(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	personnel := Personnel{
		Name:       "Poramee",
		Personalid: "1104200227093",
		//BirthDay:   time.Date(2001, 8, 2, 0, 0, 0, 0, time.UTC),
		BirthDay: time.Now().Add(24 * time.Hour), // อนาคต, fail
		Tel:      "0638267373",
		Address:  "889 หมู่ 1 ต.ระแงง อ.ศีขรภูมิ จ.สุรินทร์ 32110",
		Salary:   15000,
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(personnel)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("BirthDay must be in the past"))
}

func TestSalaryMustBePositive(t *testing.T) {
	g := NewGomegaWithT(t)
	personnel := Personnel{
		Name:       "Poramee",
		Personalid: "1104200227093",
		BirthDay:   time.Date(2001, 8, 2, 0, 0, 0, 0, time.UTC),
		Tel:        "0638267373",
		Address:    "889 หมู่ 1 ต.ระแงง อ.ศีขรภูมิ จ.สุรินทร์ 32110",
		Salary:     -50,
	}

	ok, err := govalidator.ValidateStruct(personnel)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Salary must positive"))
}
