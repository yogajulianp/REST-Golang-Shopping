package entity

import(
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id       	uint      	`form:"id" json:"id" gorm:"primaryKey"`
	Title     	string  	`form:"title" json:"title" validate:"required"`
	Image	 	string  	`form:"image" json:"image"`
	Tag	    	string		`form:"tag" json:"tag" validate:"required"`
	Description	string    	`form:"description" json:"description" validate:"required"`
	Quantity	int		    `form:"quantity" json:"quantity" validate:"required"`
	Price		float32     `form:"price" json:"price" validate:"required"`
	Cart        []Cart    	`gorm:"many2many:cart_products;"`
	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`	
}