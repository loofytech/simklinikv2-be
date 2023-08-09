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

	micro.Route("/user", func(router fiber.Router) {
		router.Get("", middleware.Auth, controllers.FindUser)
		router.Post("/signin", controllers.LoginUser)
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

	log.Fatal(app.Listen(":8000"))
}
