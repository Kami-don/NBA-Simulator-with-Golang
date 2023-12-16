package mock

import (
	"context"
	"errors"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
)

type RepositoryI interface {
	Get(ctx context.Context, id int) (*entities.Player, error)
	GetAll(ctx context.Context) ([]*entities.Player, error)
	Create(ctx context.Context, p *entities.Player) error
	Update(ctx context.Context, p *entities.Player) error
	Delete(ctx context.Context, id int) error
	GetPlayersByTeamID(ctx context.Context, teamID int) ([]*entities.Player, error)
}

type MockPlayerRepository struct {
	players map[int]*entities.Player
}

func NewMockPlayerRepository() *MockPlayerRepository {
	return &MockPlayerRepository{
		players: map[int]*entities.Player{
			11111: {
				ID:     11111,
				Name:   "Player1",
				TeamID: 111,
			},
			22222: {
				ID:     22222,
				Name:   "Player2",
				TeamID: 222,
			},
			33333: {
				ID:     33333,
				Name:   "Player3",
				TeamID: 111,
			},
		},
	}
}

func (m *MockPlayerRepository) Get(ctx context.Context, id int) (*entities.Player, error) {
	if player, exists := m.players[id]; exists {
		return player, nil
	}
	return nil, errors.New("player not found")
}

func (m *MockPlayerRepository) GetAll(ctx context.Context) ([]*entities.Player, error) {
	var players []*entities.Player
	for _, player := range m.players {
		players = append(players, player)
	}
	return players, nil
}

func (m *MockPlayerRepository) Create(ctx context.Context, p *entities.Player) error {
	if _, exists := m.players[p.ID]; exists {
		return errors.New("player already exists")
	}
	m.players[p.ID] = p
	return nil
}

func (m *MockPlayerRepository) Update(ctx context.Context, p *entities.Player) error {
	if _, exists := m.players[p.ID]; !exists {
		return errors.New("player not found")
	}
	m.players[p.ID] = p
	return nil
}

func (m *MockPlayerRepository) Delete(ctx context.Context, id int) error {
	if _, exists := m.players[id]; !exists {
		return errors.New("player not found")
	}
	delete(m.players, id)
	return nil
}

func (m *MockPlayerRepository) GetPlayersByTeamID(ctx context.Context, teamID int) ([]*entities.Player, error) {
	var players []*entities.Player
	for _, player := range m.players {
		if player.TeamID == teamID {
			players = append(players, player)
		}
	}
	return players, nil
}
