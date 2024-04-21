package service

import "github.com/makiwebdeveloper/30dayschallenge-server/pkg/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
