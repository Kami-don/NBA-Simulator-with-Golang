package game

import (
	"context"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
)

type Service struct {
	repo RepositoryI
}

func NewService(repo RepositoryI) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) GetAll(tx context.Context) ([]*entities.Game, error) {
	return s.repo.GetAll(tx)
}

func (s *Service) Get(tx context.Context, id int) (*entities.Game, error) {
	return s.repo.Get(tx, id)
}

func (s *Service) Create(tx context.Context, g *entities.Game) error {
	return s.repo.Create(tx, g)
}

func (s *Service) Update(tx context.Context, g *entities.Game) error {
	return s.repo.Update(tx, g)
}

func (s *Service) Delete(tx context.Context, id int) error {
	return s.repo.Delete(tx, id)
}
