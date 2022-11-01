package entity

import(
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       	int      `form:"id" json:"id" gorm:"primaryKey"`
	Name     	string  `form:"name" json:"name" validate:"required"`
	Email	 	string  `form:"email" json:"email" validate:"required"`
	Username	string    `form:"username" json:"username" validate:"required"`
	Password    string		`form:"password" json:"password" validate:"required"` 
	Cart 		[]CartResponse 		`json:"carts"` 
	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`	
}

type UserResponse struct {
	ID 			int 		`json:"id" form:"id"  `
	Name 		string  	`json:"name" form:"name" `
	Email	 	string  	`form:"email" json:"email" `
	Username	string   	`form:"username" json:"username" `
	Password    string		`form:"password" json:"-" `
}

func (UserResponse) TableName() string {
	return "users"	
}