package route

import (
	"github.com/gofiber/fiber/v2"
	"devtech/rest-golang-shopping/controllers"
	"devtech/rest-golang-shopping/config"
	"devtech/rest-golang-shopping/middleware"
)



func RouteInit(r *fiber.App) {
	//cartController := controllers.InitCartController()
	r.Static("/public", config.ProjectRootPath+ "/public/asset")


	r.Post("/login", controllers.AuthLogin)

	r.Post("/register", controllers.UserControllerRegister)
	r.Get("/user", middleware.Auth, controllers.UserControllerGetAll)
	r.Get("/user/:id", controllers.UserControllerGetById)
	r.Put("/user/:id", controllers.UserControllerUpdate)
	r.Delete("/user/:id", controllers.UserControllerDelete)

	r.Post("/product", controllers.ProductControllerCreate)
	r.Get("/product", controllers.ProductControllerGetAll)
	r.Get("/product/:id", controllers.ProductControllerGetById)
	r.Put("/product/:id", controllers.ProductControllerUpdate)
	r.Delete("/product/:id", controllers.ProductControllerDelete)

	r.Get("/cart", controllers.GetAllCart)
	r.Post("/cart", controllers.CreateCart)
}