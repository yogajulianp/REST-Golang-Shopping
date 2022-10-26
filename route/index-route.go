package route

import (
	"github.com/gofiber/fiber/v2"
	"devtech/rest-golang-shopping/controllers"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", controllers.UserControllerGetAll)
	r.Get("/user/:id", controllers.UserControllerGetById)
	r.Post("/user", controllers.UserControllerCreate)
	r.Put("/user/:id", controllers.UserControllerUpdate)
	r.Delete("/user/:id", controllers.UserControllerDelete)
}