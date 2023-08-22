package router

import (
	"log"
	"sim-klinikv2/controllers"
	"sim-klinikv2/middleware"

	"github.com/gofiber/fiber/v2"
)

func Server() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/signin", controllers.LoginUser)
	})

	// micro.Group("/user", middleware.Auth)
	micro.Group("/user", middleware.CheckUserRoleAdmin)
	micro.Route("/user", func(router fiber.Router) {
		router.Get("", controllers.FindUser)
		router.Post("/", controllers.CreateUserHandler)
	})

	micro.Route("/user/:userId", func(router fiber.Router) {
		router.Delete("", controllers.UserDelete)
		router.Get("", controllers.FindUserById)
		router.Patch("", controllers.UpdateUser)
	})

	micro.Route("/role", func(router fiber.Router) {
		router.Get("", controllers.FindRole)
		router.Post("/create", controllers.CreateRoleHandler)
		router.Post("/", controllers.CreateRoleHandler)
	})
	micro.Route("/role/:roleId", func(router fiber.Router) {
		router.Delete("", controllers.RoleDelete)
		router.Get("", controllers.FindRoleById)
		router.Patch("", controllers.UpdateRole)
	})

	micro.Route("/education", func(router fiber.Router) {
		router.Get("", controllers.FindEducation)
		router.Post("/create", controllers.CreateEducationHandler)
		router.Post("/", controllers.CreateEducationHandler)
	})
	micro.Route("/education/:educationId", func(router fiber.Router) {
		router.Delete("", controllers.EducationDelete)
		router.Get("", controllers.FindEducationById)
		router.Patch("", controllers.UpdateEducation)
	})

	micro.Route("/job", func(router fiber.Router) {
		router.Get("", controllers.FindJob)
		router.Post("/create", controllers.CreateJobHandler)
		router.Post("/", controllers.CreateJobHandler)
	})
	micro.Route("/job/:jobId", func(router fiber.Router) {
		router.Delete("", controllers.JobDelete)
		router.Get("", controllers.FindJobById)
		router.Patch("", controllers.UpdateJob)
	})

	micro.Route("/marital", func(router fiber.Router) {
		router.Get("", controllers.FindMaritalStatus)
		router.Post("/create", controllers.CreateMaritalStatusHandler)
		router.Post("/", controllers.CreateMaritalStatusHandler)
	})
	micro.Route("/marital/:maritalStatusId", func(router fiber.Router) {
		router.Delete("", controllers.MaritalStatusDelete)
		router.Get("", controllers.FindMaritalStatusById)
		router.Patch("", controllers.UpdateMaritalStatus)
	})

	micro.Route("/ethnic", func(router fiber.Router) {
		router.Get("", controllers.FindEthnic)
		router.Post("/create", controllers.CreateEthnicHandler)
		router.Post("/", controllers.CreateEthnicHandler)
	})
	micro.Route("/ethnic/:ethnicId", func(router fiber.Router) {
		router.Delete("", controllers.EthnicDelete)
		router.Get("", controllers.FindEthnicById)
		router.Patch("", controllers.UpdateEthnic)
	})

	micro.Route("/service", func(router fiber.Router) {
		router.Get("", controllers.FindService)
		router.Post("/create", controllers.CreateServiceHandler)
		router.Post("/", controllers.CreateServiceHandler)
	})
	micro.Route("/service/:serviceId", func(router fiber.Router) {
		router.Delete("", controllers.ServiceDelete)
		router.Get("", controllers.FindServiceById)
		router.Patch("", controllers.UpdateService)
	})

	micro.Route("/relation-agency", func(router fiber.Router) {
		router.Get("", controllers.FindRelationAgency)
		router.Post("/create", controllers.CreateRelationAgencyHandler)
		router.Post("/", controllers.CreateRelationAgencyHandler)
	})
	micro.Route("/relation-agency/:relationAgencyId", func(router fiber.Router) {
		router.Delete("", controllers.RelationAgencyDelete)
		router.Get("", controllers.FindRelationAgencyById)
		router.Patch("", controllers.UpdateRelationAgency)
	})

	micro.Route("/insurance-product", func(router fiber.Router) {
		router.Get("", controllers.FindInsuranceProduct)
		router.Post("/create", controllers.CreateInsuranceProductHandler)
		router.Post("/", controllers.CreateInsuranceProductHandler)
	})
	micro.Route("/insurance-product/:insuranceProductId", func(router fiber.Router) {
		router.Delete("", controllers.InsuranceProductDelete)
		router.Get("", controllers.FindInsuranceProductById)
		router.Patch("", controllers.UpdateInsuranceProduct)
	})

	micro.Route("/unit", func(router fiber.Router) {
		router.Get("", controllers.FindUnit)
		router.Post("/create", controllers.CreateUnitHandler)
		router.Post("/", controllers.CreateUnitHandler)
	})
	micro.Route("/unit/:unitId", func(router fiber.Router) {
		router.Delete("", controllers.UnitDelete)
		router.Get("", controllers.FindUnitById)
		router.Patch("", controllers.UpdateUnit)
	})
	micro.Route("unit/service/:serviceId", func(router fiber.Router) {
		router.Get("", controllers.FindUnitByServiceId)
	})

	micro.Route("/religion", func(router fiber.Router) {
		router.Get("", controllers.FindReligion)
		router.Post("/create", controllers.CreateReligionHandler)
		router.Post("/", controllers.CreateReligionHandler)
	})
	micro.Route("/religion/:religionId", func(router fiber.Router) {
		router.Delete("", controllers.ReligionDelete)
		router.Get("", controllers.FindReligionById)
		router.Patch("", controllers.UpdateReligion)
	})

	micro.Route("/patient", func(router fiber.Router) {
		router.Get("", controllers.FindPatient)
		router.Post("/create", controllers.CreatePatientHandler)
		router.Post("/", controllers.CreatePatientHandler)
	})
	micro.Route("/patient/:patientId", func(router fiber.Router) {
		router.Delete("", controllers.PatientDelete)
		router.Get("", controllers.FindPatientById)
		router.Patch("", controllers.UpdatePatient)
	})

	micro.Route("/registration", func(router fiber.Router) {
		router.Get("", controllers.FindRegistration)
		router.Post("/create", controllers.CreateRegistrationHandler)
		router.Post("/", controllers.CreateRegistrationHandler)
	})
	micro.Route("/registration/:registrationId", func(router fiber.Router) {
		router.Delete("", controllers.RegistrationDelete)
		router.Get("", controllers.FindRegistrationById)
		router.Patch("", controllers.UpdateRegistration)
	})

	micro.Route("/service-action", func(router fiber.Router) {
		router.Get("", controllers.FindServiceAction)
		router.Post("/create", controllers.CreateServiceActionHandler)
		router.Post("/", controllers.CreateServiceActionHandler)
	})
	micro.Route("/service-action/:serviceActionId", func(router fiber.Router) {
		router.Delete("", controllers.ServiceActionDelete)
		router.Get("", controllers.FindServiceActionById)
		router.Patch("", controllers.UpdateServiceAction)
	})

	micro.Route("/payment", func(router fiber.Router) {
		router.Get("", controllers.FindPayment)
		router.Post("/create", controllers.CreatePaymentHandler)
		router.Post("/", controllers.CreatePaymentHandler)
	})
	micro.Route("/payment/:paymentId", func(router fiber.Router) {
		router.Delete("", controllers.PaymentDelete)
		router.Get("", controllers.FindPaymentById)
		router.Patch("", controllers.UpdatePayment)
	})

	micro.Route("/doctor-schedule", func(router fiber.Router) {
		router.Get("", controllers.FindDoctorSchedule)
		router.Post("/create", controllers.CreateDoctorScheduleHandler)
		router.Post("/", controllers.CreateDoctorScheduleHandler)
	})
	micro.Route("/doctor-schedule/:doctorScheduleId", func(router fiber.Router) {
		router.Delete("", controllers.DoctorScheduleDelete)
		router.Get("", controllers.FindDoctorScheduleById)
		router.Patch("", controllers.UpdateDoctorSchedule)
	})

	micro.Route("/diagnoses", func(router fiber.Router) {
		router.Get("", controllers.FindDiagnoses)
	})
	micro.Route("/diagnoses/:diagnosesName", func(router fiber.Router) {
		router.Get("", controllers.FindDiagnosesById)
	})

	log.Fatal(app.Listen(":8000"))
}
