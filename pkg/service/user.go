package service

import (
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/domain"
	"github.com/makiwebdeveloper/30dayschallenge-server/pkg/repository"
)

type User interface {
	FindById(id int) (*domain.User, error)
}

type UserRepository struct {
	repos *repository.Repository
}

func NewUserRepository(repos *repository.Repository) *UserRepository {
	return &UserRepository{repos: repos}
}

func (s *UserRepository) FindById(id int) (*domain.User, error) {
	user, err := s.repos.User.FindById(id)
	return user, err
}
