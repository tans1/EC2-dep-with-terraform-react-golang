package blog

import (
	"log"

	"gorm.io/gorm"

	"github.com/tans1/go-web-server/domain"
	blog_dtos "github.com/tans1/go-web-server/internal/dtos/blog"
	interfaces "github.com/tans1/go-web-server/internal/interfaces/repository_interface"
	"github.com/tans1/go-web-server/repository"
	"github.com/tans1/go-web-server/schema"
	"github.com/tans1/go-web-server/utils"
)

type BlogService struct {
	repository interfaces.BlogRepository
}

func NewBlogService(db *gorm.DB) *BlogService {
	repository := repository.NewBlogRepository(db)
	return &BlogService{
		repository: repository,
	}
}

func (s *BlogService) Create(blog *blog_dtos.NewBlog, userId uint64) (*schema.Blog, error) {

	newBlog := &domain.Blog{
		Title:       blog.Title,
		Description: blog.Description,
		UserID:      userId,
	}
	createdBlog, err := s.repository.Create(newBlog)

	if err != nil {
		log.Panic("blog_service.go : ", err)
	}
	blogSchema, err := utils.TypeConverter[schema.Blog](createdBlog)
	if err != nil {
		log.Panic("blog_service.go : ", err)
	}
	return blogSchema, nil
}
func (s *BlogService) GetById(id uint64) (*schema.BlogWithComment, error) {
	commentsLimit := 5
	blog, err := s.repository.GetById(id, uint64(commentsLimit))
	if err != nil {
		log.Panic("blog_service.go : ", err)
	}

	blogSchema, err := utils.TypeConverter[schema.BlogWithComment](blog)
	comments, err2 := utils.ListTypeConverter[schema.Comment](blogSchema.Comments)
	if err != nil || err2 != nil {
		log.Panic("blog_service.go : ", err)
	}

	blogSchema.Comments = comments

	return blogSchema, nil
}
func (s *BlogService) Update(newBlog *domain.Blog) (*domain.Blog, error) {
	return nil, nil
}
func (s *BlogService) Delete(id uint64) error {
	blog, err := s.repository.GetById(id, 0)
	if err != nil {
		log.Panic("blog_service.go : ", err)
	}
	err = s.repository.Delete(*blog, id)
	if err != nil {
		log.Panic("blog_service.go : ", err)
	}
	return nil
}
