package model

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,max=60"`
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,min=8,max=60"`
	Email    string `json:"email" validate:"required,email"`
	Msisdn   string `json:"msisdn" validate:"required,msisdn"`
}
