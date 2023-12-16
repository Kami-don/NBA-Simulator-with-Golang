package simulation

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
)

type Player struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TeamID int    `json:"team_id"`
}

type Team struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Abbreviation string   `json:"abbreviation"`
	Players      []Player `json:"players"`
}

type TeamData struct {
	Teams []Team `json:"teams"`
}

func (s *Simulate) FillDBForSimulate() {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("JSON Error:", err)
		return
	}
	defer jsonFile.Close()

	var jsonData TeamData
	err = json.NewDecoder(jsonFile).Decode(&jsonData)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	data := jsonData.Teams
	for _, team := range data {
		t := &entities.Team{
			ID:           team.ID,
			Abbreviation: team.Abbreviation,
			Name:         team.Name,
		}
		s.teamService.Create(s.ctx, t)

		for _, player := range team.Players {
			p := &entities.Player{
				ID:     player.ID,
				Name:   player.Name,
				TeamID: player.TeamID,
			}
			s.playerService.Create(s.ctx, p)

		}
	}
}
