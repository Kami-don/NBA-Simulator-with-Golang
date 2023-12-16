package playergame

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

func (s *Service) Get(ctx context.Context, id int) (*entities.PlayerStats, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) GetAll(ctx context.Context) ([]*entities.PlayerStats, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) Create(ctx context.Context, p *entities.PlayerStats) error {
	return s.repo.Create(ctx, p)
}

func (s *Service) Update(ctx context.Context, p *entities.PlayerStats) error {
	return s.repo.Update(ctx, p)
}

func (s *Service) Delete(ctx context.Context, id int) error {

	return s.repo.Delete(ctx, id)
}

func (s *Service) GetAssistLeader(ctx context.Context) (*entities.PlayerStats, error) {
	return s.repo.GetAssistLeader(ctx)
}

func (s *Service) GetPlayersByGameID(ctx context.Context, teamID int) ([]*entities.PlayerStats, error) {
	return s.repo.GetPlayersByGameID(ctx, teamID)
}
