package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/tans1/go-web-server/internal/services/auth"
	"github.com/tans1/go-web-server/repository"
	"github.com/tans1/go-web-server/utils"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized,
				utils.Response{
					Success: false,
					Message: "Authorization header missing",
					Data:    nil,
				})
			ctx.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized,
				utils.Response{
					Success: false,
					Message: "Authorization header format must be Bearer {token}",
					Data:    nil,
				})
			ctx.Abort()
			return
		}

		token := parts[1]
		if token == "" || token == "null"{
			ctx.JSON(http.StatusUnauthorized,
				utils.Response{
					Success: false,
					Message: "Token is missing",
					Data:    nil,
				})
			ctx.Abort()
			return
		}
		userRepository := repository.NewUserRepository(db)
		newUserService := auth.NewAuthService(userRepository)
		user, err := newUserService.DecodeToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized,
				utils.Response{
					Success: false,
					Message: "Invalid token",
					Data:    nil,
				})
			ctx.Abort()
			return
		}
		ctx.Set("userId", user.ID)
		ctx.Next()
	}
}
