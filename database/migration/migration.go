package migration

import (
	"fmt"
	"devtech/rest-golang-shopping/database"
	"devtech/rest-golang-shopping/models/entity"
	"log"
)

func RunMigration() {
	err := database.Db.AutoMigrate(
		&entity.User{}, 
		&entity.Product{}, 
		&entity.Cart{},
		&entity.Transaction{},
	)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}