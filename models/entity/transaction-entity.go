package entity

import(
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID       	int      	`form:"id" json:"id" gorm:"primaryKey"`
	UserId     	int  		`form:"userid" json:"userid" validate:"required"`
	CartId		int  		`form:"cardid" json:"cardid" validate:"required"`
	Status		string	    `form:"status" json:"status" validate:"required"`
	User      	User		
	Product     Product		
	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`	
}

// CRUD
func AddnewTrasaction(db *gorm.DB, item *Transaction) (err error) {
	err = db.Create(item).Error
	if err != nil {
		return err
	}
	return nil
}
func GetTransaction(db *gorm.DB, transaction *[]Transaction, id int)(err error) {
	err = db.Where("userId = ? ", id).Preload("Cart").Preload("User").Find(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}
