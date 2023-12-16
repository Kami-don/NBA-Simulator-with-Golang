package service

import (
	"context"
	"fmt"
	"math"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/game"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/team"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
)

type ScoreboardService struct {
	teamService *team.Service
	gameService *game.Service
}

func NewScoreboardService(teamService *team.Service, gameService *game.Service) *ScoreboardService {
	return &ScoreboardService{
		teamService: teamService,
		gameService: gameService,
	}
}

func (t *ScoreboardService) GetScoreboard() []*entities.Scoreboard {
	data, err := t.gameService.GetAll(context.Background())
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}

	var games []*entities.Scoreboard
	for _, game := range data {
		home_team, err1 := t.teamService.Get(context.Background(), game.Home.TeamID)
		away_team, err2 := t.teamService.Get(context.Background(), game.Away.TeamID)
		fmt.Print("home_team: ", home_team)
		if err1 != nil || err2 != nil {
			continue
		}

		games = append(games, &entities.Scoreboard{
			HomeScore:       game.Home.Score,
			AwayScore:       game.Away.Score,
			HomeTeam:        home_team.Name,
			AwayTeam:        away_team.Name,
			HomeSuccessRate: math.Round(game.Home.SuccessRate * 100),
			AwaySuccessRate: math.Round(game.Away.SuccessRate * 100),
		})
	}

	return games
}
