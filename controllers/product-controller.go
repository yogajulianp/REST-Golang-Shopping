package controllers

import (
	//"strconv"
	"devtech/rest-golang-shopping/database"
	"devtech/rest-golang-shopping/models/entity"
	//"devtech/rest-golang-shopping/utils"

	//"devtech/rest-golang-shopping/models/request"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// func UserControllerGetAll(c *fiber.Ctx) error  {
// 	var users []entity.User
// 	result := database.Db.Find(&users)
// 	if result.Error != nil {
// 		log.Println(result.Error)
// 	}
// 	return c.JSON(users)
// }

func ProductControllerCreate(c *fiber.Ctx) error {
	//product := new(entity.Product)
	var product entity.Product

	if err := c.BodyParser(&product); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(product)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	// handle file
	file, errFile := c.FormFile("image")
	if errFile != nil {
		log.Println("Error File : ", errFile)
	}

	filename := file.Filename
	errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))
	if errSaveFile != nil {
		log.Println("Failed to save file into public images")
	}

	newProduct := entity.Product{
		Title   : product.Title,
		Image	: filename,
		Tag		: product.Tag,
		Description: product.Description,
	}

	
	errCreateProduct := database.Db.Create(&newProduct).Error
	if errCreateProduct  != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal menyimpan data",
		})
	}
	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success",
		"data": newProduct,
	})
}