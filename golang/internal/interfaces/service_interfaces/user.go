package interfaces

import "github.com/tans1/go-web-server/internal/dtos/user"

type User interface {
	FindById(id int) (*user_dto.User, error)
	DeleteAccount(id int) error
}
