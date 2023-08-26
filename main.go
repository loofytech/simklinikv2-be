package main

import (
	"log"

	"sim-klinikv2/config"
	"sim-klinikv2/router"
)

func init() {
	cfs, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	config.ConnectDB(&cfs)

	// var patient *models.Patient
	// var count int64
	// config.DB.Find(&patient).Count(&count)
	// // mr := strings.Join([]string{"A", fmt.Sprintf("%06d", record.RowsAffected+1)}, "")
	// // fmt.Print(mr)
	// fmt.Print(count)
}

func main() {
	router.Server()
	// micro := fiber.New()

	// app.Mount("/api", micro)
	// app.Use(logger.New())

	// micro.Route("/user", func(router fiber.Router) {
	// 	router.Get("", middleware.Auth, controllers.FindUser)
	// 	router.Post("/signin", controllers.LoginUser)
	// 	router.Post("/", controllers.CreateUserHandler)
	// 	router.Get("/test", middleware.Auth, func(c *fiber.Ctx) error {
	// 		return c.SendString("hello")
	// 	})
	// })
	// // app.Use(cors.New(cors.Config{
	// // 	AllowOrigins:     "http://localhost:3000",
	// // 	AllowHeaders:     "Origin, Content-Type, Accept",
	// // 	AllowMethods:     "GET, POST, PATCH, DELETE",
	// // 	AllowCredentials: true,
	// // }))

	// // jwt := middleware.NewAuthMiddleware("supersecretkey")

	// // Checker api live
	// micro.Get("/healthchecker", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{
	// 		"status":  "success",
	// 		"message": "Welcome to Golang, Fiber, and GORM",
	// 	})
	// })
}
