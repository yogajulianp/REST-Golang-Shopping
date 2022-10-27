package route

import (
	"github.com/gofiber/fiber/v2"
	"devtech/rest-golang-shopping/controllers"
	"devtech/rest-golang-shopping/config"
	"devtech/rest-golang-shopping/middleware"
)



func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectRootPath+ "/public/asset")

	r.Post("/login", controllers.Login)

	r.Get("/user", middleware.Auth, controllers.UserControllerGetAll)
	r.Get("/user/:id", controllers.UserControllerGetById)
	r.Post("/user", controllers.UserControllerCreate)
	r.Put("/user/:id", controllers.UserControllerUpdate)
	r.Delete("/user/:id", controllers.UserControllerDelete)

	r.Post("/product", controllers.ProductControllerCreate)
}