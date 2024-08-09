package interfaces

import (
	"github.com/gin-gonic/gin"

	comment_dto "github.com/tans1/go-web-server/internal/dtos/comment"
	"github.com/tans1/go-web-server/schema"
)

type (
	CommentController interface {
		Create(ctx *gin.Context)
		GetById(ctx *gin.Context)
	}
	CommentService interface {
		Create(comment *comment_dto.CreateComment, userId uint64, blogId uint64) (*schema.Comment, error)
		GetById(commentId uint64, userId uint64) (*schema.CommentWithReply, error)
	}
)
