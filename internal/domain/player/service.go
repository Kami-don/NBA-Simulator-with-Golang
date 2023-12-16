package player

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

func (s *Service) Get(ctx context.Context, id string) (*entities.Player, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*entities.Player, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, p *entities.Player) error {
	return s.repo.Create(ctx, p)
}

func (s *Service) Update(ctx context.Context, p *entities.Player) error {
	return s.repo.Update(ctx, p)
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
