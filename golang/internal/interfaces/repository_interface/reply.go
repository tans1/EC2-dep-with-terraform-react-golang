package interfaces

import "github.com/tans1/go-web-server/domain"

type ReplyRepository interface {
	GenericRepository[domain.Reply]

	GetById(id uint64) (*domain.Reply, error)
	GetRepliesByCommentId(id uint64, limit uint64) ([]*domain.Reply, error)
}
