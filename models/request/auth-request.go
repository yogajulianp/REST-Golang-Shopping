package request


type LoginRequest struct {
	Username	string    `form:"username" json:"username" validate:"required"`
	Password    string	`form:"password" json:"password" validate:"required"`
}