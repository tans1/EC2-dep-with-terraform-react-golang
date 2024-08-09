package repository

import (
	"gorm.io/gorm"

	"github.com/tans1/go-web-server/domain"
)

type ReplyRepository struct {
	GenericRepository[domain.Reply]
	db *gorm.DB
}

func NewReplyRepository(db *gorm.DB) *ReplyRepository {
	return &ReplyRepository{
		db:                db,
		GenericRepository: GenericRepository[domain.Reply]{db: db},
	}
}

func (r *ReplyRepository) GetById(id uint64) (*domain.Reply, error) {
	var reply domain.Reply
	result := r.db.Preload("CreatedBy").First(&reply, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &reply, nil
}

func (r *ReplyRepository) GetRepliesByCommentId(id uint64, limit uint64) ([]*domain.Reply, error) {
	var replies []*domain.Reply

	result := r.db.Preload("CreatedBy").Preload("Blog").Where("comment_id = ?", id).Limit(int(limit)).Find(&replies)
	if result.Error != nil {
		return nil, result.Error
	}

	return replies, nil
}
