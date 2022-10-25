package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UserControllerRead(c *fiber.Ctx) error  {
		return c.JSON(fiber.Map{
			"data": "user",
		})
	
}