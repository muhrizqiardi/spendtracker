package service

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhrizqiardi/spendtracker/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	LogIn(payload dto.LogInDTO) (string, error)
}

type authService struct {
	us     UserService
	secret string
}

func NewAuthService(us UserService, secret string) *authService {
	return &authService{us, secret}
}

func (as *authService) LogIn(payload dto.LogInDTO) (string, error) {
	user, err := as.us.GetOneByEmail(payload.Email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return "", err
	}

	userIDStr := strconv.Itoa(int(user.ID))
	claim := jwt.RegisteredClaims{
		Subject:   userIDStr,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(as.secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}
