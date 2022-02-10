package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBillPass(t *testing.T) {
	g := NewGomegaWithT(t)

	bill := Bill{
		Cot:        125,
		Com:        125,
		Sum:        250,
		Listofbill: 2,
		Dateofbill: time.Date(2023, 2, 13, 0, 0, 0, 0, time.UTC),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bill)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())

}

func TestBillSumMustnotzero(t *testing.T) {
	g := NewGomegaWithT(t)

	bill := Bill{
		Cot:        125.22,
		Com:        125.22,
		Sum:        0.00, //false
		Listofbill: 2,
		Dateofbill: time.Date(2023, 2, 13, 0, 0, 0, 0, time.UTC),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bill)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Sum must not Zero"))

}
func TestBillListmustnotzero(t *testing.T) {
	g := NewGomegaWithT(t)

	bill := Bill{
		Cot:        125.22,
		Com:        125.22,
		Sum:        250.44,
		Listofbill: 0, //false
		Dateofbill: time.Date(2023, 2, 13, 0, 0, 0, 0, time.UTC),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bill)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("List must not Zero"))

}

func TestBillDate(t *testing.T) {
	g := NewGomegaWithT(t)

	bill := Bill{
		Cot:        125.22,
		Com:        125.22,
		Sum:        250.44,
		Listofbill: 2,
		Dateofbill: time.Date(2011, 9, 30, 22, 0, 0, 0, time.Local), //false
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(bill)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ไม่สามารถบันทึกเป็นเวลาในอดีตได้"))

}
