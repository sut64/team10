package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAppointmentPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	appointment := Appointment{
		Appoint_ID:   "A1234",
		Room_number:  1,
		Date_appoint: time.Date(2022, time.Month(3), 10, 10, 0, 0, 0, &time.Location{}),
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(appointment)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestAppointIDMustBeValidPattern(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"D0000",
		"AA000",
		"A00000",
		"A000A",
	}

	for _, fixture := range fixtures {
		appointment := Appointment{
			Appoint_ID:   fixture,
			Room_number:  5,
			Date_appoint: time.Date(2022, time.Month(3), 11, 10, 0, 0, 0, &time.Location{}),
		}

		ok, err := govalidator.ValidateStruct(appointment)

		// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
		g.Expect(ok).ToNot(BeTrue())

		// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
		g.Expect(err).ToNot(BeNil())

		// err.Error ต้องมี error message แสดงออกมา
		g.Expect(err.Error()).To(Equal("Appoint ID must Start with A and (0-9) 4 digits"))
	}
}

func TestDateAppointMustbeFuture(t *testing.T) {
	g := NewGomegaWithT(t)

	appointment := Appointment{
		Appoint_ID:   "A4321",
		Room_number:  3,
		Date_appoint: time.Date(2022, time.Month(1), 13, 10, 0, 0, 0, &time.Location{}),
	}
	ok, err := govalidator.ValidateStruct(appointment)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Date Appointment must be in the future"))

}

func TestRoomNumberMustGreaterZero(t *testing.T) {
	g := NewGomegaWithT(t)

	appointment := Appointment{
		Appoint_ID:   "A4321",
		Room_number:  -2,
		Date_appoint: time.Date(2022, time.Month(3), 12, 10, 0, 0, 0, &time.Location{}),
	}
	ok, err := govalidator.ValidateStruct(appointment)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(Equal("Room number greater than zero value"))

}
