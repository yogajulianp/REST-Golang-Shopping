package entity

import(
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        	uint     `gorm:"primaryKey"`
	Id       	int     `form:"id" json:"id" validate:"required"`
	Name     	string  `form:"name" json:"name" validate:"required"`
	Email	 	string  `form:"email" json:"email" validate:"required"`
	Username	string    `form:"username" json:"username" validate:"required"`
	Password    string	`form:"password" json:"password" validate:"required"`
	CreatedAt time.Time		`json:"created_at"`
  	UpdatedAt time.Time		`json:"updated_at"`
 	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	
}