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
		Name:       "Poramee Suriyajanno",
		Personalid: "1104200227093",
		BirthDay:   time.Date(2001, 8, 2, 0, 0, 0, 0, time.UTC),
		Tel:        "0638267373",
		Address:    "889 หมู่ 1 ต.ระแงง อ.ศีขรภูมิ จ.สุรินทร์ 32110",
		Salary:     15000,
	}

	ok, err := govalidator.ValidateStruct(personnel)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())

}

func TestBirthDayMustBePast(t *testing.T) {
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

func TestSalaryMustBePositive(t *testing.T) {
	g := NewGomegaWithT(t)
	personnel := Personnel{
		Name:       "Poramee Suriyajanno",
		Personalid: "1104200227093",
		BirthDay:   time.Date(2001, 8, 2, 0, 0, 0, 0, time.UTC),
		Tel:        "0638267373",
		Address:    "889 หมู่ 1 ต.ระแงง อ.ศีขรภูมิ จ.สุรินทร์ 32110",
		Salary:     15000,
	}

	ok, err := govalidator.ValidateStruct(personnel)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())

}
