package interfaces

import "github.com/tans1/go-web-server/domain"

type BlogRepository interface {
	GenericRepository[domain.Blog] // will have all the methods of the generic repository,

	GetById(id uint64, limit uint64) (*domain.Blog, error)
	Update(newBlog *domain.Blog) (*domain.Blog, error)
}
