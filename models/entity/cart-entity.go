package entity

import(
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	ID       	int      	`form:"id" json:"id" gorm:"primaryKey"`
	Quantity	int		    `form:"quantity" json:"quantity" validate:"required"`
	Total		float32     `form:"total" json:"total" `
	Status		string	    `form:"status" json:"status" `
	UserID     	int  		`form:"user_id" json:"user_id" validate:"required"`
	User 		UserResponse 	`json:"users"`
	Products 	[]Product 	`json:"products" gorm:"many2many:cart_products"`
	ProductID	int  		`form:"product_id" json:"product_id" gorm:"-"`

	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`	
}

type CartResponse struct {
	ID       	int      	`form:"id" json:"id" `
	Quantity	int		    `form:"quantity" json:"quantity" `
	Total		float32     `form:"total" json:"total" `
	Status		string	    `form:"status" json:"status" `
	UserID     	int  		`form:"user_id" json:"user_id" `
	ProductID	int  		`form:"product_id" json:"product_id" `
	User		 UserResponse 	`json:"users"`
	Product		[]Product		`json:"products" gorm:"many2many:cart_products;ForeignKey:ID;joinForeignKey:CartID;References:ID;joinReferences:ProductID"`
}

type CartResponseTransaction struct {
	ID       	int      	`form:"id" json:"id" `
	Quantity	int		    `form:"quantity" json:"quantity" `
	Total		float32     `form:"total" json:"total" `
	Status		string	    `form:"status" json:"status" `
}


func (CartResponse) TableName() string {
	return "carts"	
}

func (CartResponseTransaction) TableName() string {
	return "carts"	
}



