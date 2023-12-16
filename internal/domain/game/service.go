package game

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

func (s *Service) GetAll(ctx context.Context) ([]*entities.Game, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Get(ctx context.Context, id int) (*entities.Game, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) Create(ctx context.Context, g *entities.Game) error {
	return s.repo.Create(ctx, g)
}

func (s *Service) Update(ctx context.Context, g *entities.Game) error {
	return s.repo.Update(ctx, g)
}

func (s *Service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
