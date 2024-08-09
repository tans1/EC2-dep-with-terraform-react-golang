package user

import user_dto "github.com/tans1/go-web-server/internal/dtos/user"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (uServ *UserService) FindById(id int) (*user_dto.User, error) {
	return &user_dto.User{
		FirstName: "",
		LastName:  "",
	}, nil
}

func (uServ *UserService) DeleteAccount(id int) error {
	return nil
}
