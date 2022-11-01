package entity

import(
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID       	int      	`form:"id" json:"id" gorm:"primaryKey"`
	Title     	string  	`form:"title" json:"title" validate:"required"`
	Image	 	string  	`form:"image" json:"image"`
	Tag	    	string		`form:"tag" json:"tag" validate:"required"`
	Description	string    	`form:"description" json:"description" validate:"required"`
	Quantity	int		    `form:"quantity" json:"quantity" validate:"required"`
	Price		float32     `form:"price" json:"price" validate:"required"`
		
	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`	
}

type ProductResponse struct {
	ID       	int      	`form:"id" json:"id" `
	Title     	string  	`form:"title" json:"title"`
	Image	 	string  	`form:"image" json:"image"`
	Tag	    	string		`form:"tag" json:"tag" `
	Description	string    	`form:"description" json:"description" `
	Quantity	int		    `form:"quantity" json:"quantity" `
	Price		float32     `form:"price" json:"price" `
}

func (ProductResponse) TableName() string {
	return "products"	
}