package users

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	signUp(ctx *gin.Context, body *SignUpReq) error
}
type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) signUp(ctx *gin.Context, body *SignUpReq) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		slog.Error("failed to hash password", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return err
	}

	user := &User{Email: body.Email, Password: string(hash)}
	err = s.repo.create(user)
	if err != nil {
		return err
	}

	return nil
}
