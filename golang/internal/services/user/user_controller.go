package user

import (
	"github.com/tans1/go-web-server/internal/dtos/user"
	"github.com/tans1/go-web-server/internal/interfaces/service_interfaces"
)

type UserController struct {
	service interfaces.User
}

func NewUserController() *UserController {
	service := NewUserService()
	return &UserController{
		service: service,
	}
}

func (uCont *UserController) FindById(id int) (*user_dto.User, error) {
	return uCont.service.FindById(id)
}

func (uCont *UserController) DeleteAccount(id int) error {
	return uCont.service.DeleteAccount(id)
}
