package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestHistorySheetPass(t *testing.T) {
	g := NewGomegaWithT(t)

	historysheet := HistorySheet{
		Weight:      50.00,
		Height:      150.59,
		Temperature: 35.25,
		PressureOn:  100,
		PressureLow: 80,
		Symptom:     "มีไข้ ไอ เจ็บคอ",
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(historysheet)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())

}
func TestHistorySheetWeight(t *testing.T) {
	g := NewGomegaWithT(t)

	historysheet := HistorySheet{
		Weight:      -0.9,
		Height:      150.52,
		Temperature: 35.5,
		PressureOn:  100,
		PressureLow: 80,
		Symptom:     "มีไข้ ไอ เจ็บคอ",
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(historysheet)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Weight must be positive"))

}
