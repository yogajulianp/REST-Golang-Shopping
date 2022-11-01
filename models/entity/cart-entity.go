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
	User		 UserResponse 	`json:"user"`
	Product		[]Product		`json:"products" gorm:"many2many:cart_products;ForeignKey:ID;joinForeignKey:CartID;References:ID;joinReferences:ProductID"`
}

func (CartResponse) TableName() string {
	return "carts"	
}



// // CRUD
// func AddnewCart(db *gorm.DB, item *Cart) (err error) {
// 	err = db.Create(item).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func GetCartbyUser(db *gorm.DB, cart *[]Cart, id int)(err error) {
// 	err = db.Where(&Cart {UserId: id, Status: "process"}).Preload("User").Preload("Product").Find(&cart).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func ListCartbyId(db *gorm.DB, cart *Cart, productid int, userid int)(err error) {
// 	err = db.Where(&Cart {ProductId: productid, UserId: userid, Status: "process"}).Preload("User").Preload("Product").First(&cart).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func UpdateCart(db *gorm.DB, cart *Cart)(err error) {
// 	db.Save(cart)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func DeleteCartById(db *gorm.DB, item *Cart, id int)(err error) {
// 	db.Where("id=?", id).Delete(item)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }