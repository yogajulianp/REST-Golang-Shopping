package main

import (
	"github.com/gofiber/fiber/v2"
	"devtech/rest-golang-shopping/route"
	"devtech/rest-golang-shopping/database"
	"devtech/rest-golang-shopping/database/migration"
)

func main() {

	// initial database
	database.InitDb()
	migration.RunMigration()

	app := fiber.New()

	//initial route
	route.RouteInit(app)

	app.Get("/", func(c  *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})


	app.Listen(":3000")
}