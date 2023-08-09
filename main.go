package main

import (
	"log"

	"sim-klinikv2/config"
	"sim-klinikv2/controllers"

	"sim-klinikv2/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	cfs, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	config.ConnectDB(&cfs)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/api", micro)
	app.Use(logger.New())
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://localhost:3000",
	// 	AllowHeaders:     "Origin, Content-Type, Accept",
	// 	AllowMethods:     "GET, POST, PATCH, DELETE",
	// 	AllowCredentials: true,
	// }))

	// jwt := middleware.NewAuthMiddleware("supersecretkey")

	micro.Route("/user", func(router fiber.Router) {
		router.Get("", middleware.Auth, controllers.FindUser)
		router.Post("/signin", controllers.LoginUser)
		router.Post("/", controllers.CreateUserHandler)
		router.Get("/test", middleware.Auth, func(c *fiber.Ctx) error {
			return c.SendString("hello")
		})
	})
	micro.Route("/user/:userId", func(router fiber.Router) {
		router.Delete("", controllers.UserDelete)
		router.Get("", controllers.FindUserById)
		router.Patch("", controllers.UpdateUser)
	})

	micro.Route("/role", func(router fiber.Router) {
		router.Get("", controllers.FindUser)
		router.Post("/create", controllers.CreateRoleHandler)
		router.Post("/", controllers.CreateUserHandler)
	})
	micro.Route("/role/:roleId", func(router fiber.Router) {
		router.Delete("", controllers.UserDelete)
		router.Get("", controllers.FindUserById)
		router.Patch("", controllers.UpdateRole)
	})

	// Checker api live
	micro.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
