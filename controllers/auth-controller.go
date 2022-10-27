package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"devtech/rest-golang-shopping/models/request"
	"devtech/rest-golang-shopping/models/entity"
	"devtech/rest-golang-shopping/utils"
	"devtech/rest-golang-shopping/database"
	"fmt"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := c.BodyParser(&loginRequest); err != nil {
		return err
	}
	fmt.Println(loginRequest)

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	//check available user
	var user entity.User
	err := database.Db.Where("username = ?", loginRequest.Username).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	//check validasi password
	validPassword := utils.CheckPasswordHash(loginRequest.Password, user.Password) 
	if !validPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token": "secret",
	})
}
