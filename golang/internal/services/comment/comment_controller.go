package comment

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	comment_dto "github.com/tans1/go-web-server/internal/dtos/comment"
	interfaces "github.com/tans1/go-web-server/internal/interfaces/service_interfaces"
	"github.com/tans1/go-web-server/utils"
)

type CommentController struct {
	service interfaces.CommentService
}

func NewCommentController(db *gorm.DB) *CommentController {
	service := NewCommentService(db)
	return &CommentController{
		service: service,
	}
}

func (c *CommentController) Create(ctx *gin.Context) {
	var comment comment_dto.CreateComment
	if err := ctx.BindJSON(&comment); err != nil {
		ctx.JSON(400, utils.Response{
			Success: false,
			Message: "Invalid data",
			Data:    nil,
		})
		return
	}
	userId, ok := ctx.Get("userId")
	if !ok {
		log.Panic("not able to get the userId from the context")
	}
	blogId := ctx.Param("id")
	blogId64, err := strconv.ParseUint(blogId, 10, 32)
	if err != nil {
		log.Panic(err)
	}
	result, err := c.service.Create(&comment, userId.(uint64), blogId64)
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
		Message: "Comment created successfully",
		Data: map[string]interface{}{
			"comment": result,
		},
	})

}
func (c *CommentController) GetById(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		log.Panic("not able to get the userId from the context")
	}
	commentId := ctx.Param("id")
	commentId64, err := strconv.ParseUint(commentId, 10, 32)
	if err != nil {
		log.Panic(err)
	}
	result, err := c.service.GetById(commentId64, userId.(uint64))
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
		Message: "Comment found",
		Data: map[string]interface{}{
			"comment": result,
		},
	})

}
