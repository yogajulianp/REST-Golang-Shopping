package controllers

import (
	"devtech/rest-golang-shopping/database"
	"devtech/rest-golang-shopping/models/entity"
	//"devtech/rest-golang-shopping/utils"

	//"devtech/rest-golang-shopping/models/request"
	
	"log"
	

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


func CreateCart(c *fiber.Ctx) error {
	carts := new(entity.Cart)
	if err := c.BodyParser(&carts); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(carts)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	if carts.ProductID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"err": "product_id is required",
		})
	}

	var product entity.Product
	newCart := entity.Cart{
		Quantity : carts.Quantity,
		Total : float32(carts.Quantity)*product.Price,
		Status: "process",
		UserID: carts.UserID,
		ProductID: carts.ProductID,
	}

	errCreatePost := database.Db.Debug().Create(&newCart).Error
	if errCreatePost != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal melakukan order cart" + errCreatePost.Error(),
		})
	}

		cartProduct := new(entity.CartProduct)
		cartProduct.CartID = newCart.ID
		cartProduct.ProductID = newCart.ProductID
		database.Db.Create(&cartProduct)
	

	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success order carts",
		"carts": newCart,
	})
}

func GetAllCart(c *fiber.Ctx) error  {
	var carts []entity.CartResponse
	result := database.Db.Preload("User").Preload("Product").Find(&carts)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(fiber.Map{
		"carts": carts,
	})
}
