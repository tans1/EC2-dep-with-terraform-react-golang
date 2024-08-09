package repository

import (
	"gorm.io/gorm"

	"github.com/tans1/go-web-server/domain"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) Create(comment *domain.Comment) (*domain.Comment, error) {
	result := r.db.Create(comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}
func (r *CommentRepository) GetById(commentId uint64, repliesLimit uint64) (*domain.Comment, error) {
	var comment domain.Comment
	result := r.db.Preload("CreatedBy").
		Preload("Blog").
		Preload("Reply", func(db *gorm.DB) *gorm.DB {
			return db.Limit(int(repliesLimit))
		}).
		First(&comment, commentId)

	if result.Error != nil {
		return nil, result.Error
	}
	return &comment, nil
}
