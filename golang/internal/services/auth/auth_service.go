package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/tans1/go-web-server/domain"
	auth_dtos "github.com/tans1/go-web-server/internal/dtos/auth"
	"github.com/tans1/go-web-server/repository"
	"github.com/tans1/go-web-server/schema"
	"github.com/tans1/go-web-server/utils"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(repository *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: repository,
	}
}

func (c *AuthService) Register(authDto *auth_dtos.AuthSignUp) (*schema.User, error) {
	hashedPassword, err := hashPassword(authDto.Password)
	if err != nil {
		log.Panic("auth_service.go : ", err)
	}
	user, err := utils.TypeConverter[domain.User](authDto)
	if err != nil {
		log.Panic(err)
	}
	user.Password = hashedPassword

	newUser, err := c.userRepository.Create(user)
	if err != nil {
		log.Panic("auth_service.go : ", err)
	}
	userResp, err := utils.TypeConverter[schema.User](newUser)
	if err != nil {
		log.Panic("auth_service.go : ", err)
	}
	return userResp, nil
}

func (c *AuthService) Login(authDto *auth_dtos.AuthLogin) (*schema.User, error) {
	user, err := c.userRepository.GetByUsername(authDto.Username)
	correctPassword := checkPasswordHash(authDto.Password, user.Password)

	if err != nil || correctPassword == false {
		return nil, utils.ErrorResponse{
			StatusCode: 400,
			Message:    "Invalid Credential",
		}
	}
	
	userResp, err := utils.TypeConverter[schema.User](user)
	if err != nil {
		log.Panic("auth_service.go", err)
	}
	return userResp, nil
}

func (c *AuthService) GenerateToken(user *schema.User) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"iss": "golang-project",
		"aud": "user",
		"exp": time.Now().Add(3 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	secretKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		log.Panic("auth_setting.go : ", err)
	}

	return tokenString
}

func (c *AuthService) DecodeToken(tokenString string) (*schema.User, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, utils.ErrorResponse{
				StatusCode: 400,
				Message:    "Unexpected signing method",
			}
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Panic("auth_setting.go : ", err)
	}
	if !token.Valid {
		return nil, utils.ErrorResponse{
			StatusCode: 400,
			Message:    "Invalid Token",
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, utils.ErrorResponse{
			StatusCode: 400,
			Message:    "Invalid Claims",
		}
	}

	username := claims["sub"].(string)
	user, err := c.userRepository.GetByUsername(username)
	if err != nil {
		log.Panic("auth_setting.go : ", err)
	}

	userResp, err := utils.TypeConverter[schema.User](user)
	if err != nil {
		log.Panic("auth_setting.go : ", err)
	}

	return userResp, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
