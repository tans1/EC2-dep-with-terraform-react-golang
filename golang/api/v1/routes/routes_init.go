package routes

import (
	"gorm.io/gorm"

	"github.com/tans1/go-web-server/internal/services/auth"
	"github.com/tans1/go-web-server/internal/services/blog"
	"github.com/tans1/go-web-server/internal/services/comment"
	"github.com/tans1/go-web-server/internal/services/reply"
)

type Controllers struct {
	authController    *auth.AuthController
	blogController    *blog.BlogController
	commentController *comment.CommentController
	replyController *reply.ReplyController
}

func New(db *gorm.DB) *Controllers {
	return &Controllers{
		authController:    auth.NewAuthController(db),
		blogController:    blog.NewBlogController(db),
		commentController: comment.NewCommentController(db),
		replyController: reply.NewReplyController(db),
	}
}
