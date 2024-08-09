package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/tans1/go-web-server/domain"
)

type BlogRepository struct {
	GenericRepository[domain.Blog]
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		GenericRepository: GenericRepository[domain.Blog]{db: db},
		db:                db,
	}
}

func (r *BlogRepository) GetById(id uint64, limit uint64) (*domain.Blog, error) {
	var blog domain.Blog
	result := r.db.Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Limit(int(limit))
	}).Find(&blog, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blog, nil
}

func (r *BlogRepository) Update(newBlog *domain.Blog) (*domain.Blog, error) {
	var blog domain.Blog
	result := r.db.First(&blog, newBlog.ID)
	if result.Error != nil {
		return nil, result.Error
	}

	if blog.ID != newBlog.ID {
		return nil, errors.New("not author of the blog")
	}
	blog.Title = newBlog.Title
	blog.Description = newBlog.Description

	r.db.Save(&blog)
	return &blog, nil
}
