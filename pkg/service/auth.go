package service

import (
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/domain"
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/model"
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	SignUp(body *model.SignUpRequest) error
	SignIn(body *model.SignInRequest) (*domain.User, error)
}

type AuthService struct {
	repos *repository.Repository
}

func NewAuthService(repos *repository.Repository) *AuthService {
	return &AuthService{repos: repos}
}

func (s *AuthService) SignUp(body *model.SignUpRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return err
	}

	user := domain.User{
		Username: body.Username,
		Password: string(hash),
		Name:     body.Name,
	}

	s.repos.User.Create(&user)

	return nil
}

func (s *AuthService) SignIn(body *model.SignInRequest) (*domain.User, error) {
	user, err := s.repos.User.FindByUsername(body.Username)

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, err
	}

	return user, nil
}
