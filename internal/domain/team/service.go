package team

import (
	"context"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
)

type Service struct {
	repo RepositoryI
}

func NewService(repo RepositoryI) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetAll(tx context.Context) ([]*entities.Team, error) {
	return s.repo.GetAll(tx)
}

func (s *Service) Get(tx context.Context, id string) (*entities.Team, error) {
	return s.repo.Get(tx, id)
}

func (s *Service) Create(tx context.Context, t *entities.Team) error {
	return s.repo.Create(tx, t)
}

func (s *Service) Update(tx context.Context, t *entities.Team) error {
	return s.repo.Update(tx, t)
}

func (s *Service) Delete(tx context.Context, id string) error {
	return s.repo.Delete(tx, id)
}
