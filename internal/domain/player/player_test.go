package player

import (
	"context"
	"testing"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/infra/mock"
	"github.com/stretchr/testify/assert"
)

func TestPlayerUnit(t *testing.T) {
	mockPlayerRepo := mock.NewMockPlayerRepository()
	mockPlayerService := NewService(mockPlayerRepo)

	t.Run("Get", func(t *testing.T) {
		player, err := mockPlayerService.Get(context.Background(), "11111")
		assert.Equal(t, player.Name, "Player1")
		assert.Equal(t, player.Team, "TeamA")

		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("GetAll", func(t *testing.T) {
		players, err := mockPlayerService.GetAll(context.Background())
		assert.Equal(t, len(players), 3)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("Create", func(t *testing.T) {
		err := mockPlayerService.Create(context.Background(), &entities.Player{
			ID:   "44444",
			Name: "Player4",
			Team: "TeamD",
		})
		players, _ := mockPlayerService.GetAll(context.Background())
		assert.Equal(t, len(players), 4)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		err := mockPlayerService.Update(context.Background(), &entities.Player{
			ID:   "44444",
			Name: "Player4",
			Team: "TeamE",
		})
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := mockPlayerService.Delete(context.Background(), "44444")
		players, _ := mockPlayerService.GetAll(context.Background())
		assert.Equal(t, len(players), 3)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})
}
