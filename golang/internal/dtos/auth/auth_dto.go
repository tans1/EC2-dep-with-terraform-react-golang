package auth_dtos

type AuthSignUp struct {
	Username    string `json:"username" validate:"required,min=5"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,password"` // the word is custom tag of the custom validation function
	FirstName   string `json:"firstName" validate:"required,min=10"`
	LastName    string `json:"lastName" validate:"required,min=10"`
	PhoneNumber string `json:"phone" validate:"required,min=10,max=15"`
}

type AuthLogin struct {
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required"`
}
