package service

import "github.com/makiwebdeveloper/30dayschallenge-server/pkg/repository"

type Service struct {
	Auth
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos),
		User: NewUserRepository(repos),
	}
}
