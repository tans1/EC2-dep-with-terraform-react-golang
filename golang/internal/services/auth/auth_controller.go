package auth

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"

	"github.com/tans1/go-web-server/internal/dtos"
	auth_dtos "github.com/tans1/go-web-server/internal/dtos/auth"
	interfaces "github.com/tans1/go-web-server/internal/interfaces/service_interfaces"
	"github.com/tans1/go-web-server/repository"
	"github.com/tans1/go-web-server/utils"
)

type AuthController struct {
	service interfaces.AuthService
}

func NewAuthController(db *gorm.DB) *AuthController {
	repository := repository.NewUserRepository(db)
	service := NewAuthService(repository)
	return &AuthController{
		service: service,
	}

}

func (c *AuthController) Register(ctx *gin.Context) {

	validate := validator.New()
	validate.RegisterValidation("password", auth_dtos.PasswordValidator)

	var user auth_dtos.AuthSignUp
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(400, utils.Response{
			Success: false,
			Message: "Invalid data",
			Data:    nil,
		})
		return
	}

	if err := validate.Struct(&user); err != nil {
		validationErrors := dtos.TranslateValidationError(err)
		ctx.JSON(400, utils.Response{
			Success: false,
			Message: "Invalid data provided",
			Data:    validationErrors,
		})
		return
	}

	if _, err := c.service.Register(&user); err != nil {
		statusCode, message := utils.DecodeError(err)
		ctx.JSON(statusCode, utils.Response{
			Success: false,
			Message: message,
			Data:    nil,
		})
	} else {
		ctx.JSON(200, utils.Response{
			Success: true,
			Message: "user created",
			Data:    nil,
		})
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	validate := validator.New()

	var userCred auth_dtos.AuthLogin
	if err := ctx.BindJSON(&userCred); err != nil {
		ctx.JSON(400, utils.Response{
			Success: false,
			Message: "Invalid data",
			Data:    nil,
		})
		return
	}
	// validate credentials
	if err := validate.Struct(&userCred); err != nil {
		validationErrors := dtos.TranslateValidationError(err)
		ctx.JSON(400, utils.Response{
			Success: false,
			Message: "Invalid data provided",
			Data:    validationErrors,
		})

		return
	}

	// authenticate user
	if user, err := c.service.Login(&userCred); err != nil {
		statusCode, message := utils.DecodeError(err)
		ctx.JSON(statusCode, utils.Response{
			Success: false,
			Message: message,
			Data:    nil,
		})
	} else {
		token := c.service.GenerateToken(user)
		ctx.JSON(200, utils.Response{
			Success: true,
			Message: "user logged in",
			Data: map[string]interface{}{
				"token": token,
				"user":  *user,
			},
		})
	}
}
