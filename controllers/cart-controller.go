package controllers

import (
	"gorm.io/gorm"
	//"strconv"
	"devtech/rest-golang-shopping/database"
	"devtech/rest-golang-shopping/models/entity"
	//"devtech/rest-golang-shopping/utils"

	//"devtech/rest-golang-shopping/models/request"
	"strconv"
	// "fmt"
	

	//"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CartController struct {
	// declare variables
	Db *gorm.DB
}

func InitCartController() *CartController {
	db := database.InitDb()
	// gorm
	db.AutoMigrate(&entity.Cart{})

	return &CartController{Db: db}
}

func (controller *CartController) CartControllerGet(c *fiber.Ctx) error  {
	userid := c.Query("userid")
	userId,_ := strconv.Atoi(userid)

	var cart []entity.Cart
	errResult := entity.GetCartbyUser(controller.Db, &cart, userId)
	if errResult != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "server error, tidak bisa get cart",
		})
	}
	return c.JSON(cart)
}

func (controller *CartController) CartControllerRequestOrder(c *fiber.Ctx) error {
	userid := c.Query("userid")
	userId,_ := strconv.Atoi(userid)

	productid := c.Query("productid")
	productId,_ := strconv.Atoi(productid)

	
	var product entity.Product
	var cart entity.Cart
	var addCart entity.Cart

	//get product by id
	errGetProductId := database.Db.Where("id = ?", productId).First(&product).Error
	if errGetProductId != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	errCart := entity.ListCartbyId(controller.Db, &cart, productId, userId)
	if errCart != nil {
		return c.Status(500).JSON(fiber.Map{
			"message" : "server error, tidak muncul di list",
		})
	}

	if cart.Id != 0 {
		cart.Quantity = cart.Quantity + 1
		cart.Total = cart.Total + product.Price

		errUpdate := database.Db.Save(&cart).Error
		if errUpdate != nil {
			return c.Status(500).JSON(fiber.Map{
				"message" : "server error",
			})
		}
		//if succeed
		return c.JSON(fiber.Map{
			"message"	: "success",
			"data"		:    cart,
		})
	} else {
		addCart.UserId = userId
		addCart.ProductId = productId
		addCart.Quantity = 1
		cart.Total = float32(cart.Quantity)*product.Price

		errAddCart := database.Db.Create(&addCart).Error
		if errAddCart != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal menyimpan data",
		})
		}
		//if succeed
		return c.JSON(fiber.Map{
			"message": "success",
			"data":    addCart,
		})

	}

}
