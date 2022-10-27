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

func ProductControllerGetAll(c *fiber.Ctx) error  {
	var products []entity.Product
	result := database.Db.Find(&products)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(products)
}

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
		Quantity: product.Quantity,
		Price : product.Price,
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

func ProductControllerGetById(c *fiber.Ctx) error {
	productId := c.Params("id")
	
	fmt.Println(productId)

	var product entity.Product
	err := database.Db.Where("id = ?", productId).First(&product).Error
	
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	
	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success",
		"data": product,
	})
}


func ProductControllerUpdate(c *fiber.Ctx) error {
	productId := c.Params("id")

	var product entity.Product
	err := database.Db.Where("id = ?", productId).First(&product).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	var productRequest entity.Product
	if err := c.BodyParser(&productRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
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

	var filename string = file.Filename
	errSaveFile := c.SaveFile(file, fmt.Sprintf("./public/images/%s", filename))
	if errSaveFile != nil {
		log.Println("Failed to save file into public images")
	}

	
	
	//Update User Data  
	product.Title = productRequest.Title
	productRequest.Image = filename
	product.Image = productRequest.Image
	product.Tag = productRequest.Tag
	product.Description = productRequest.Description
	product.Quantity = productRequest.Quantity
	product.Price = productRequest.Price

	errUpdate := database.Db.Save(&product).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "server error",
		})
	}

	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success data has been updated",
		"data": product,
	})

}

func ProductControllerDelete(c *fiber.Ctx) error {
	productId := c.Params("id")

	var product entity.Product
	err := database.Db.Debug().Where("id = ?", productId).First(&product).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.Db.Debug().Delete(&product).Error
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