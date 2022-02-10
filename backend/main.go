package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sut64/team10/controller"
	"github.com/sut64/team10/entity"
)

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())
	// User Routes
	r.GET("/users", controller.ListUsers)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	// BloodType Routes of Patientrecord
	r.GET("/bloodtypes", controller.ListBloodType)
	r.GET("/bloodtype/:id", controller.GetBloodType)
	r.POST("/bloodtypes", controller.CreateBloodType)
	r.PATCH("/bloodtypes", controller.UpdateBloodType)
	r.DELETE("/bloodtypes/:id", controller.DeleteBloodType)

	// Gender Routes of Patientrecord
	r.GET("/genders", controller.ListGender)
	r.GET("/gender/:id", controller.GetGender)
	r.POST("/genders", controller.CreateGender)
	r.PATCH("/genders", controller.UpdateGender)
	r.DELETE("/genders/:id", controller.DeleteGender)

	// Personnel Routes of Patientrecord
	r.GET("/personnels", controller.ListPersonnel)
	r.GET("/personnel/:id", controller.GetPersonnel)
	r.POST("/personnels", controller.CreatePersonnel)
	r.PATCH("/personnels", controller.UpdatePersonnel)
	r.DELETE("/personnels/:id", controller.DeletePersonnel)

	// Personnel Routes of PatientrecordCreate
	r.GET("/personnelPatientRecords", controller.ListPersonnelPatientRecord)

	// Prename Routes of Patientrecord
	r.GET("/prenames", controller.ListPrenames)
	r.GET("/prename/:id", controller.GetPrename)
	r.POST("/prenames", controller.CreatePrename)
	r.PATCH("/prenames", controller.UpdatePrename)
	r.DELETE("/prenames/:id", controller.DeletePrename)

	// Province Routes of Patientrecord
	r.GET("/provinces", controller.ListProvinces)
	r.GET("/province/:id", controller.GetProvince)
	r.POST("/provinces", controller.CreateProvince)
	r.PATCH("/provinces", controller.UpdateProvince)
	r.DELETE("/provinces/:id", controller.DeleteProvince)

	// Patientrecord Routes
	r.GET("/patientrecords", controller.ListPatientrecords)
	r.GET("/patientrecord/:id", controller.GetPatientrecord)
	r.POST("/patientrecords", controller.CreatePatientrecord)
	r.PATCH("/patientrecords", controller.UpdatePatientrecord)
	r.DELETE("/patientrecords/:id", controller.DeletePatientrecord)

	// Appointment Routes
	r.GET("/appointments", controller.ListAppointments)         //list
	r.GET("/appointment/:id", controller.GetAppointment)        //get
	r.POST("/appointments", controller.CreateAppointment)       //create
	r.PATCH("/appointments", controller.UpdateAppointment)      //update
	r.DELETE("/appointments/:id", controller.DeleteAppointment) //delete

	// DrugAllergy Routes
	r.GET("/drugallergys", controller.ListDrugAllergys)
	r.GET("/drugallergy/:id", controller.GetDrugAllergy)
	r.POST("/drugallergys", controller.CreateDrugAllergy)
	r.PATCH("/drugallergys", controller.UpdateDrugAllergy)
	r.DELETE("/drugallergys/:id", controller.DeleteDrugAllergy)

	// HistorySheet Routes
	r.GET("/historysheets", controller.ListHistorySheets)
	r.GET("/historysheet/:id", controller.GetHistorySheet)
	r.POST("/historysheets", controller.CreateHistorySheet)
	r.PATCH("/historysheets", controller.UpdateHistorySheet)
	r.DELETE("/historysheets/:id", controller.DeleteHistorySheet)

	// medicine Routes
	r.GET("/medicine", controller.ListMedicine)
	r.GET("/medicine/:id", controller.GetMedicne)
	r.POST("/medicine", controller.CreateMedicine)
	r.PATCH("/medicine", controller.UpdateMedicine)
	r.DELETE("/medicine/:id", controller.DeleteMedicine)

	// medicaltreatment Routes
	r.GET("/medicaltreatment", controller.ListMedicalTreatment)
	r.GET("/medicaltreatment/:id", controller.GetMedicalTreatment)
	r.POST("/medicaltreatment", controller.CreateMedicalTreatment)
	r.PATCH("/medicaltreatment", controller.UpdateMedicalTreatment)
	r.DELETE("/medicaltreatment/:id", controller.DeleteMedicalTreatment)

	// bill Routes
	r.GET("/bill", controller.ListBill)
	r.GET("/bill/:id", controller.GetBill)
	r.POST("/bill", controller.CreateBill)
	r.PATCH("/bill", controller.UpdateBill)
	r.DELETE("/bill/:id", controller.DeleteBill)

	// treatmentrecord Routes
	r.GET("/treatmentrecord", controller.ListTreatment)
	r.GET("/treatmentrecord/:id", controller.GetTreatment)
	r.POST("/treatmentrecord", controller.CreateTreatment)
	r.PATCH("/treatmentrecord", controller.UpdateTreatment)
	r.DELETE("/treatmentrecord/:id", controller.DeleteTreatment)

	// disease Routes
	r.GET("/disease", controller.ListDisease)
	r.GET("/disease/:id", controller.GetDisease)
	r.POST("/disease", controller.CreateDisease)
	r.PATCH("/disease", controller.UpdateDisease)
	r.DELETE("/disease/:id", controller.DeleteDisease)

	//jobtitle routes
	r.GET("/jobtitle", controller.ListJobTitle)
	r.GET("/jobtitle/:id", controller.GetJobTitle)
	r.POST("/jobtitle", controller.CreateJobTitle)
	r.PATCH("/jobtitle", controller.UpdateJobTitle)
	r.DELETE("/jobtitle/:id", controller.DeleteJobTitle)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}

}
