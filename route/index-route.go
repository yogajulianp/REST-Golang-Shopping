package route

import (
	"github.com/gofiber/fiber/v2"
	"devtech/rest-golang-shopping/controllers"
)

func RouteInit(r *fiber.App) {
	r.Get("/user", controllers.UserControllerRead)
}