package controllers

import (
	//"strconv"
	"devtech/rest-golang-shopping/database"
	"devtech/rest-golang-shopping/models/entity"
	//"devtech/rest-golang-shopping/models/request"
	"github.com/go-playground/validator/v10"
	"log"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func UserControllerGetAll(c *fiber.Ctx) error  {
	var users []entity.User
	result := database.Db.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(users)
}

func UserControllerCreate(c *fiber.Ctx) error {
	user := new(entity.User)
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}
	newUser := entity.User{
		Name    : user.Name,
		Email	: user.Email,
		Username : user.Username,
		Password : user.Password,
	}
	errCreateUser := database.Db.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal menyimpan data",
		})
	}
	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success",
		"data": newUser,
	})
}

func UserControllerGetById(c *fiber.Ctx) error {
	userId := c.Params("id")
	//userId,_ := strconv.Atoi(Id)
	fmt.Println(userId)

	var user entity.User
	err := database.Db.Where("id = ?", userId).First(&user).Error
	
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	
	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success",
		"data": user,
	})

}

func UserControllerUpdate(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entity.User
	err := database.Db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	userRequest := new(entity.User)
	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	//Update User Data  
	if userRequest.Username != "" {
		user.Username = userRequest.Username
	}
	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Password = userRequest.Password

	errUpdate := database.Db.Save(&user).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "server error",
		})
	}

	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success data has been updated",
		"data": user,
	})

}

func UserControllerDelete(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entity.User
	err := database.Db.Debug().Where("id = ?", userId).First(&user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.Db.Debug().Delete(&user).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "server error",
		})
	}

		//if succeed
		return c.JSON(fiber.Map{
			"message" : "user has been deleted",
		})
}