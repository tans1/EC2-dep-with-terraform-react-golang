package interfaces

import "github.com/tans1/go-web-server/domain"

type CommentRepository interface {
	Create(comment *domain.Comment) (*domain.Comment, error)
	GetById(commentId uint64, repliesLimit uint64) (*domain.Comment, error)
}
