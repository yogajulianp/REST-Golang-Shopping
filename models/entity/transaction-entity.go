package entity

import(
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID       	int      	`form:"id" json:"id" gorm:"primaryKey"`
	Status		string	    `form:"status" json:"status" `
	UserID     	int  		`form:"user_id" json:"user_id" validate:"required"`
	User      	[]UserResponse		`json:"users" gorm:"many2many:transaction_users"`
	Cart   	[]CartResponse	`json:"carts" gorm:"many2many:transaction_carts"`
	CartID		int  		`form:"cart_id" json:"cart_id" gorm:"-"`	

	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`	
}

type TransactionResponse struct {
	ID       	int      	`form:"id" json:"id"`
	UserID     	int  		`form:"user_id" json:"user_id" `
	CartID		int  		`form:"cart_id" json:"cart_id" `
	Status		string	    `form:"status" json:"status" `
	User      	[]UserResponse		`json:"users" gorm:"many2many:transaction_users;ForeignKey:ID;joinForeignKey:TransactionID;References:ID;joinReferences:UserID" `
	Cart   		[]CartResponse	`json:"carts" gorm:"many2many:transaction_carts;ForeignKey:ID;joinForeignKey:TransactionID;References:ID;joinReferences:CartID"`	
}

func (TransactionResponse) TableName() string {
	return "transactions"	
}

type TransactionCart struct {
	TransactionID int `json:"transaction_id"`
	CartID int `json:"cart_id"`
}

