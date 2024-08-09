package interfaces

import (
	"github.com/tans1/go-web-server/domain"
)

type UserRepository interface {
	// Create(user *domain.User) (*domain.User, error)
	GenericRepository[domain.User]
	GetById(id int) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}
