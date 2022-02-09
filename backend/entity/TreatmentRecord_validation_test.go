package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestTreatmentRecordPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	treatment := TreatmentRecord{
		Treatment:   "วิธีการรักษา",
		Temperature: 32,
		RecordDate:  time.Date(2022, 12, 22, 0, 0, 0, 0, time.UTC),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(treatment)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestDateMustBeFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	treatment := TreatmentRecord{
		Treatment:   "วิธีการรักษา",
		Temperature: 32,
		RecordDate:  time.Date(2021, 12, 22, 0, 0, 0, 0, time.UTC),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(treatment)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("ไม่สามารถบันทึกเป็นเวลาในอดีตได้"))
}


