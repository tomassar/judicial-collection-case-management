package auth

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/lawyers"
	"github.com/tomassar/judicial-collection-case-management/internal/domain/users"
	"golang.org/x/crypto/bcrypt"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

type Service interface {
	SignUp(ctx *gin.Context, body *SignUpReq) error
	Login(ctx *gin.Context, body *LoginReq) error
	Logout(ctx *gin.Context)
}

type service struct {
	userService   users.Service
	lawyerService lawyers.Service
}

func NewService(userService users.Service, lawyerService lawyers.Service) Service {
	return &service{
		userService:   userService,
		lawyerService: lawyerService,
	}
}

func (s *service) SignUp(ctx *gin.Context, body *SignUpReq) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		slog.Error("failed to hash password", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return err
	}

	user := &users.User{Email: body.Email, Password: string(hash)}
	err = s.userService.Create(ctx, user)
	if err != nil {
		return err
	}

	err = s.lawyerService.CreateLawyer(ctx, &lawyers.CreateLawyerReq{
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Login(ctx *gin.Context, body *LoginReq) error {
	user, err := s.userService.GetByEmail(ctx, body.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * 30 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return err
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	return nil
}

func (s *service) Logout(ctx *gin.Context) {
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
}

func NewAuth() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, "http://localhost:3000/auth/google"),
	)

}
