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


func CreateTransaction(c *fiber.Ctx) error {
	transactions := new(entity.Transaction)
	if err := c.BodyParser(&transactions); err != nil {
		return err
	}

	//validasi request
	validate := validator.New()
	errValidate := validate.Struct(transactions)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error": errValidate.Error(),
		})
	}

	if transactions.CartID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"err": "cart_id is required",
		})
	}

	
	newTransaction := entity.Transaction{
		Status: "Transaksi diterima Lunas, Pesanan dikirim",
		UserID: transactions.UserID,
		CartID: transactions.CartID,
	}

	errCreatePost := database.Db.Debug().Create(&newTransaction).Error
	if errCreatePost != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "gagal melakukan transaksi" + errCreatePost.Error(),
		})
	}
	
		transactionCart := new(entity.TransactionCart)
		transactionCart.TransactionID = newTransaction.ID
		transactionCart.CartID = newTransaction.CartID
		database.Db.Debug().Create(&transactionCart)
	


	//if succeed
	return c.JSON(fiber.Map{
		"message" : "success transaction",
		"transaction": newTransaction,
	})
}

func GetAllTransaction(c *fiber.Ctx) error  {
	var transactions []entity.Transaction
	result := database.Db.Debug().Preload("User.Username").Preload("Cart.Product").Find(&transactions)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(fiber.Map{
		"transaction": transactions,
	})
}
