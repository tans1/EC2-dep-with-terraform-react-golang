package reply

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	reply_dto "github.com/tans1/go-web-server/internal/dtos/reply"
	interfaces "github.com/tans1/go-web-server/internal/interfaces/service_interfaces"
	"github.com/tans1/go-web-server/utils"
)

type ReplyController struct {
	service interfaces.ReplyService
}

func NewReplyController(db *gorm.DB) *ReplyController {
	service := NewReplyService(db)
	return &ReplyController{
		service: service,
	}
}

func (r *ReplyController) Create(ctx *gin.Context) {
	var replyData reply_dto.ReplyCreateDto
	if err := ctx.BindJSON(&replyData); err != nil {
		ctx.JSON(400, utils.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	userId, ok := ctx.Get("userId")
	if !ok {
		log.Panic("un able to get the userId")
	}
	commentId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Panic(err)
	}

	result, err := r.service.Create(&replyData, userId.(uint64), uint64(commentId))
	if err != nil {
		statusCode, message := utils.DecodeError(err)
		ctx.JSON(statusCode, utils.Response{
			Success: false,
			Message: message,
			Data:    nil,
		})
		return
	}
	ctx.JSON(200, utils.Response{
		Success: true,
		Message: "Created Successfully",
		Data: map[string]interface{}{
			"reply": result,
		},
	})

}

func (r *ReplyController) GetById(ctx *gin.Context) {
	replyId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Panic(err)
	}

	result, err := r.service.GetById(uint64(replyId))
	if err != nil {
		statusCode, message := utils.DecodeError(err)
		ctx.JSON(statusCode, utils.Response{
			Success: false,
			Message: message,
			Data:    nil,
		})
		return
	}
	ctx.JSON(200, utils.Response{
		Success: true,
		Message: "Data found",
		Data: map[string]interface{}{
			"reply": result,
		},
	})
}

func (r *ReplyController) GetRepliesByCommentId(ctx *gin.Context) {
	commentId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		log.Panic(err)
	}

	result, err := r.service.GetRepliesByCommentId(uint64(commentId))
	if err != nil {
		statusCode, message := utils.DecodeError(err)
		ctx.JSON(statusCode, utils.Response{
			Success: false,
			Message: message,
			Data:    nil,
		})
		return
	}
	ctx.JSON(200, utils.Response{
		Success: true,
		Message: "Data found",
		Data: result,
		
	})
}
