package interfaces

import (
	"github.com/gin-gonic/gin"

	"github.com/tans1/go-web-server/domain"
	blog_dtos "github.com/tans1/go-web-server/internal/dtos/blog"
	"github.com/tans1/go-web-server/schema"
)

type (
	BlogController interface {
		Create(ctx *gin.Context)
		GetById(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
	}

	BlogService interface {
		Create(blog *blog_dtos.NewBlog, userId uint64) (*schema.Blog, error)
		GetById(id uint64) (*schema.BlogWithComment, error)
		Update(newBlog *domain.Blog) (*domain.Blog, error)
		Delete(id uint64) error
	}
)
