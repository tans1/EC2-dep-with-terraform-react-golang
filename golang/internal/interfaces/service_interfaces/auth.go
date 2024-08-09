package interfaces

import (
	"github.com/gin-gonic/gin"

	auth_dtos "github.com/tans1/go-web-server/internal/dtos/auth"
	"github.com/tans1/go-web-server/schema"
)

type (
	AuthController interface {
		Register(ctx *gin.Context) // handler function
		Login(ctx *gin.Context)    // handler function
	}

	AuthService interface {
		Register(authDto *auth_dtos.AuthSignUp) (*schema.User, error)
		Login(loginDto *auth_dtos.AuthLogin) (*schema.User, error)
		GenerateToken(user *schema.User) string
		DecodeToken(token string) (*schema.User, error)
	}
)
