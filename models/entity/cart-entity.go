package entity

import(
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	Id       	uint      	`form:"id" json:"id" gorm:"primaryKey"`
	UserId     	int  		`form:"userid" json:"userid" validate:"required"`
	ProductId	int  		`form:"productid" json:"productid" validate:"required"`
	Quantity	int		    `form:"quantity" json:"quantity" validate:"required"`
	Total		float32     `form:"total" json:"total" validate:"required"`
	Status		string	    `form:"status" json:"status" validate:"required"`
	Product 	Product 	`gorm:"foreignkey:ProductId;references:Id"`
	User 		User 		`gorm:"foreignkey:UserId;references:Id"`
	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`	
}

// CRUD
func AddnewCart(db *gorm.DB, item *Cart) (err error) {
	err = db.Create(item).Error
	if err != nil {
		return err
	}
	return nil
}
func GetCartbyUser(db *gorm.DB, cart *[]Cart, id int)(err error) {
	err = db.Where(&Cart {UserId: id, Status: "process"}).Preload("User").Preload("Product").Find(&cart).Error
	if err != nil {
		return err
	}
	return nil
}
func ListCartbyId(db *gorm.DB, cart *Cart, productid int, userid int)(err error) {
	err = db.Where(&Cart {ProductId: productid, UserId: userid, Status: "process"}).Preload("User").Preload("Product").First(&cart).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateCart(db *gorm.DB, cart *Cart)(err error) {
	db.Save(cart)
	if err != nil {
		return err
	}
	return nil
}
func DeleteCartById(db *gorm.DB, item *Cart, id int)(err error) {
	db.Where("id=?", id).Delete(item)
	if err != nil {
		return err
	}
	return nil
}