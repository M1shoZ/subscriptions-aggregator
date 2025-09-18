package main

import (
	"log"

	"github.com/M1shoZ/subscriptions-aggregator/database"
	"github.com/M1shoZ/subscriptions-aggregator/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/M1shoZ/subscriptions-aggregator/docs"
)

// @title           Subscription API
// @version         1.0
// @description     API for managing users and subscriptions.
// @termsOfService  http://example.com/terms/

// @contact.name   M1shoZ Dev
// @contact.email  m1shoz@example.com

// @host      localhost:3000
// @BasePath  /
func main() {
	log.Println("Запускаем сервер на порте 3000...")
	database.ConnectDb()

	app := fiber.New()

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

	setupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Get("/home", handlers.Home)
	app.Get("/get_users", handlers.GetAllUsers)
	app.Post("/create_sub", handlers.CreateSub)
	app.Get("/get_subs", handlers.GetAllSubs)
	app.Get("/get_subs/:id", handlers.GetSubById)
	app.Get("/get_subs_by_user/:user_id", handlers.GetSubByUserId)
	app.Delete("/delete/:id", handlers.DeleteSub)
	app.Patch("/update/:id", handlers.UpdateSub)

	app.Get("/get_sum/:user_id/:service_name/:start_date/:end_date", handlers.GetSum)
}
