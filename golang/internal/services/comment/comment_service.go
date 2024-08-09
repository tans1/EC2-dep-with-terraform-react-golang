package comment

import (
	"log"

	"gorm.io/gorm"

	"github.com/tans1/go-web-server/domain"
	comment_dto "github.com/tans1/go-web-server/internal/dtos/comment"
	interfaces "github.com/tans1/go-web-server/internal/interfaces/repository_interface"
	"github.com/tans1/go-web-server/repository"
	"github.com/tans1/go-web-server/schema"
	"github.com/tans1/go-web-server/utils"
)

type CommentService struct {
	repository     interfaces.CommentRepository
	blogRepository interfaces.BlogRepository
}

func NewCommentService(db *gorm.DB) *CommentService {
	commentRepository := repository.NewCommentRepository(db)
	blogRepository := repository.NewBlogRepository(db)
	return &CommentService{
		repository:     commentRepository,
		blogRepository: blogRepository,
	}
}

func (s *CommentService) Create(comment *comment_dto.CreateComment, userId uint64, blogId uint64) (*schema.Comment, error) {

	blog, err := s.blogRepository.GetById(blogId, 0)
	if err != nil {
		return nil, utils.ErrorResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}
	}

	if blog == nil {
		return nil, utils.ErrorResponse{
			StatusCode: 400,
			Message:    "No blog found",
		}
	}

	commentDomain := &domain.Comment{
		Description: comment.Description,
		BlogID:      blogId,
		UserID:      userId,
	}
	result, err := s.repository.Create(commentDomain)
	
	if err != nil {
		return nil, utils.ErrorResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}
	}
	commentResp, err := utils.TypeConverter[schema.Comment](result)
	if err != nil {
		log.Panic(err)
	}

	return commentResp, nil
}
func (s *CommentService) GetById(commentId uint64, userId uint64) (*schema.CommentWithReply, error) {
	repliesLimit := 5
	comment, err := s.repository.GetById(commentId, uint64(repliesLimit))
	if err != nil {
		log.Panic(err)
	}
	commentResp, err := utils.TypeConverter[schema.CommentWithReply](comment)
	if err != nil {
		log.Panic(err)
	}

	creator, err := utils.TypeConverter[schema.User](comment.CreatedBy)
	if err != nil {
		log.Panic(err)
	}
	replies, err := utils.ListTypeConverter[schema.Reply](comment.Reply)
	if err != nil {
		log.Panic(err)
	}

	commentResp.CreatedBy = *creator
	commentResp.Replies = replies

	return commentResp, nil
}
