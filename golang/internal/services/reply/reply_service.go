package reply

import (
	"log"

	"gorm.io/gorm"

	"github.com/tans1/go-web-server/domain"
	reply_dto "github.com/tans1/go-web-server/internal/dtos/reply"
	interfaces "github.com/tans1/go-web-server/internal/interfaces/repository_interface"
	"github.com/tans1/go-web-server/repository"
	"github.com/tans1/go-web-server/schema"
	"github.com/tans1/go-web-server/utils"
)

type ReplyService struct {
	repository interfaces.ReplyRepository
}

func NewReplyService(db *gorm.DB) *ReplyService {
	repository := repository.NewReplyRepository(db)
	return &ReplyService{
		repository: repository,
	}
}

func (s *ReplyService) Create(replyDto *reply_dto.ReplyCreateDto, userId uint64, commentId uint64) (*schema.Reply, error) {
	data := &domain.Reply{
		Description: replyDto.Description,
		UserID:      userId,
		CommentID:   commentId,
		BlogID: replyDto.BlogID,
	}
	result, err := s.repository.Create(data)
	if err != nil {
		log.Panic(err)
	}

	response, err := utils.TypeConverter[schema.Reply](result)
	if err != nil {
		log.Panic(err)
	}

	return response, nil

}

func (s *ReplyService) GetById(replyId uint64) (*schema.Reply, error) {
	result, err := s.repository.GetById(replyId)
	if err != nil {
		log.Panic(err)
	}
	response, err := utils.TypeConverter[schema.Reply](result)
	if err != nil {
		log.Panic(err)
	}

	createdBy, err := utils.TypeConverter[schema.User](result.CreatedBy)
	if err != nil {
		log.Panic(err)
	}

	response.CreatedBy = *createdBy
	return response, nil
}

func (s *ReplyService) GetRepliesByCommentId(commentId uint64) (*schema.RepliesWthComment, error) {
	limit := 5
	results, err := s.repository.GetRepliesByCommentId(commentId, uint64(limit))
	if err != nil {
		log.Panic(err)
	}

	if len(results) == 0 {
		return nil, nil
	}
	response, err := utils.TypeConverter[schema.RepliesWthComment](results[0])
	replies := []schema.Reply{}

	for _, result := range results {
		if err != nil {
			log.Panic(err)
		}

		createdBy, err := utils.TypeConverter[schema.User](&result.CreatedBy)
		if err != nil {
			log.Panic(err)
		}

		reply, err := utils.TypeConverter[schema.Reply](&result)
		if err != nil {
			log.Panic(err)
		}

		reply.CreatedBy = *createdBy
		replies = append(replies, *reply)
	}

	response.Replies = replies

	return response, nil

}
