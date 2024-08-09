package interfaces

import (
	"github.com/gin-gonic/gin"

	reply_dto "github.com/tans1/go-web-server/internal/dtos/reply"
	"github.com/tans1/go-web-server/schema"
)

type (
	ReplyController interface {
		Create(ctx *gin.Context)
		GetById(ctx *gin.Context)
		GetRepliesByCommentId(ctx *gin.Context)
	}

	ReplyService interface {
		Create(replyDto *reply_dto.ReplyCreateDto, userId uint64, commentId uint64) (*schema.Reply, error)
		GetById(replyId uint64) (*schema.Reply, error)
		GetRepliesByCommentId(commentId uint64) (*schema.RepliesWthComment, error)
	}
)
