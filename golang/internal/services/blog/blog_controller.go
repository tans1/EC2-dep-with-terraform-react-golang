package blog

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	blog_dtos "github.com/tans1/go-web-server/internal/dtos/blog"
	interfaces "github.com/tans1/go-web-server/internal/interfaces/service_interfaces"
	"github.com/tans1/go-web-server/utils"
)

type BlogController struct {
	service interfaces.BlogService
}

func NewBlogController(db *gorm.DB) *BlogController {
	service := NewBlogService(db)
	return &BlogController{
		service: service,
	}
}

func (c *BlogController) Create(ctx *gin.Context) {
	validate := validator.New()

	var blog blog_dtos.NewBlog
	if err := ctx.BindJSON(&blog); err != nil {
		log.Panic("blog_controller.go", err)
	}

	if err := validate.Struct(&blog); err != nil {
		ctx.JSON(400, utils.Response{
			Success: false,
			Message: "Invalid data",
			Data:    nil,
		})
		return
	}

	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(400, utils.Response{
			Success: false,
			Message: "Author not found",
			Data:    nil,
		})
		return
	}

	result, err := c.service.Create(&blog, userId.(uint64))
	if err != nil {
		statusCode, message := utils.DecodeError(err)
		ctx.JSON(statusCode, utils.Response{
			Success: false,
			Message: message,
			Data:    nil,
		})
	} else {
		ctx.JSON(200, utils.Response{
			Success: true,
			Message: "Successfully created",
			Data: map[string]interface{}{
				"blog": result,
			},
		})
	}

}
func (c *BlogController) GetById(ctx *gin.Context) {
	blogId := ctx.Param("id")
	blogIdInt, err := strconv.ParseUint(blogId, 10, 32)
	if err != nil {
		log.Panic(err)
	}

	result, err := c.service.GetById(uint64(blogIdInt))
	if err != nil {
		statusCode, message := utils.DecodeError(err)
		ctx.JSON(statusCode, utils.Response{
			Success: false,
			Message: message,
			Data:    nil,
		})
	} else {
		ctx.JSON(200, utils.Response{
			Success: true,
			Message: "Successfully found",
			Data: map[string]interface{}{
				"blog": result,
			},
		})
	}
}
func (c *BlogController) Update(ctx *gin.Context) {
	log.Panic("not implemented")
}
func (c *BlogController) Delete(ctx *gin.Context) {
	log.Panic("not implemented")
}
