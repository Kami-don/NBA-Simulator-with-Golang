package simulation

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/game"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/player"
	playergame "github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/player_game"
	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/domain/team"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
	database "github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/infra/db"
)

type Simulate struct {
	ctx               context.Context
	db                *database.Database
	gameService       game.Service
	playerService     player.Service
	teamService       team.Service
	playerGameService playergame.Service
}

func NewSimulate(ctx context.Context, db *database.Database, gameService game.Service, playerService player.Service, teamService team.Service, playerGameService playergame.Service) *Simulate {
	return &Simulate{
		db:                db,
		gameService:       gameService,
		playerService:     playerService,
		teamService:       teamService,
		playerGameService: playerGameService,
	}
}

type MatchInfo struct {
	ID          int
	StartTime   time.Time
	ElapsedTime time.Duration
	AttackCount int
	SuccessRate float64
	IsHome      bool
}

func (s *Simulate) simulateMatch(wg *sync.WaitGroup, matchChan chan<- MatchInfo, gameID int) {
	defer wg.Done()
	startTime := time.Now()
	duration := 240 * time.Second

	matchInfo := MatchInfo{
		ID:          gameID,
		StartTime:   startTime,
		ElapsedTime: 0,
		SuccessRate: 0.0,
		IsHome:      false,
	}

	game, err := s.gameService.Get(s.ctx, gameID)
	if err != nil {
		log.Fatal("Error while getting game from database: ", err)
		return
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		isNewAttack := rand.Intn(10)
		isAssist := rand.Intn(5)
		isHomeTeam := rand.Intn(2)
		isBasket := rand.Intn(10)
		score := 0

		// 80% attack rate and 20% basket rate
		if isBasket == 0 || isNewAttack <= 2 {
			continue
		}

		if isHomeTeam == 0 {
			matchInfo.IsHome = true
		} else {
			matchInfo.IsHome = false
		}

		score = rand.Intn(2) + 2 // Get a random score of 2 or 3 points
		s.UpdatePlayersGameDB(matchInfo, score)
		if isAssist != 0 {
			s.UpdatePlayersGameDB(matchInfo, 0)
		}

		matchInfo.SuccessRate = rand.Float64()
		matchInfo.ElapsedTime = time.Since(startTime)

		if matchInfo.ElapsedTime >= duration {
			matchChan <- matchInfo
			game.IsFinished = true
			s.gameService.Update(s.ctx, game)
			return
		} else {
			s.UpdateGameDB(matchInfo, game, score)
		}

		matchChan <- matchInfo
	}

}

func (s *Simulate) UpdateGameDB(match MatchInfo, game *entities.Game, score int) {
	if match.IsHome {
		game.Home.Score += score
		game.Home.AttackCount += match.AttackCount
		game.Home.SuccessRate = match.SuccessRate
	} else {
		game.Away.Score += score
		game.Away.AttackCount += match.AttackCount
		game.Away.SuccessRate = match.SuccessRate
	}
	s.gameService.Update(s.ctx, game)
}

func (s *Simulate) UpdatePlayersGameDB(match MatchInfo, score int) {
	playersGame, err := s.playerGameService.GetPlayersByGameID(s.ctx, match.ID)
	if err != nil {
		log.Fatal("Error while getting player game from database: ", err)
		return
	}

	randomPlayer := rand.Intn(5)
	playerGame := playersGame[randomPlayer]

	if score == 2 {
		playerGame.TwoPointScore += score
	} else if score == 3 {
		playerGame.ThreePointScore += score
	} else {
		playerGame.Assist += 1
	}

	s.playerGameService.Update(s.ctx, playerGame)
}

func (s *Simulate) ChoosePlayers(team1 int, team2 int, gameID int) {
	team1Players, err := s.playerService.GetPlayersByTeamID(s.ctx, team1)
	if err != nil {
		log.Fatal("Error while getting players from database: ", err)
		return
	}

	team1GamePlayers := make([]int, 5)
	for _, player := range team1Players {
		playerGame := &entities.PlayerStats{
			ID:              rand.Intn(1000000),
			PlayerID:        player.ID,
			GameID:          gameID,
			TwoPointScore:   0,
			ThreePointScore: 0,
			Assist:          0,
		}
		s.playerGameService.Create(s.ctx, playerGame)
		team1GamePlayers = append(team1GamePlayers, playerGame.ID)
		if len(team1GamePlayers) == 5 {
			break
		}

	}

	team2Players, err := s.playerService.GetPlayersByTeamID(s.ctx, team2)
	if err != nil {
		log.Fatal("Error while getting players from database: ", err)
		return
	}

	team2GamePlayers := make([]int, 5)
	for _, player := range team2Players {
		playerGame := &entities.PlayerStats{
			ID:              rand.Intn(1000000),
			PlayerID:        player.ID,
			GameID:          gameID,
			TwoPointScore:   0,
			ThreePointScore: 0,
			Assist:          0,
		}
		s.playerGameService.Create(s.ctx, playerGame)
		team2GamePlayers = append(team2GamePlayers, playerGame.ID)
		if len(team2GamePlayers) == 5 {
			break
		}
	}
}

func (s *Simulate) GenerateRandomGames() []int {
	rand.Seed(time.Now().UnixNano())

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	matches := make(map[int]int)
	remaining := make([]int, len(numbers))

	for i, num := range numbers {
		remaining[i] = num
	}

	for len(remaining) > 1 {
		idx1 := rand.Intn(len(remaining))
		num1 := remaining[idx1]
		remaining = append(remaining[:idx1], remaining[idx1+1:]...)

		idx2 := rand.Intn(len(remaining))
		num2 := remaining[idx2]
		remaining = append(remaining[:idx2], remaining[idx2+1:]...)

		matches[num1] = num2
	}

	if len(remaining) == 1 {
		matches[remaining[0]] = remaining[0]
	}

	var gameIDs []int
	for team1, team2 := range matches {
		randomGameID := rand.Intn(1000000)
		gameIDs = append(gameIDs, randomGameID)
		s.gameService.Create(s.ctx, &entities.Game{
			ID:          randomGameID,
			StartTime:   time.Now(),
			ElapsedTime: 0,
			IsFinished:  false,
			Away: entities.TeamGameInfo{
				TeamID:      team1,
				Score:       0,
				AttackCount: 0,
				SuccessRate: 0.0,
			},
			Home: entities.TeamGameInfo{
				TeamID:      team2,
				Score:       0,
				AttackCount: 0,
				SuccessRate: 0.0,
			},
		})
		s.ChoosePlayers(team1, team2, randomGameID)
	}

	return gameIDs
}

func (s *Simulate) Run() {
	rand.Seed(time.Now().UnixNano())

	s.FillDBForSimulate()

	gameIDs := s.GenerateRandomGames()

	numMatches := len(gameIDs) // 5 match simulation
	var wg sync.WaitGroup
	matchChan := make(chan MatchInfo, numMatches)

	// Start a specified number of match simulations
	for _, gameID := range gameIDs {
		wg.Add(1)
		go s.simulateMatch(&wg, matchChan, gameID)
	}

	go func() {
		wg.Wait()
		close(matchChan)
	}()

	for match := range matchChan {
		fmt.Printf("Match ID: %d\n", match.ID)
		fmt.Printf("Elapsed Time: %.0f seconds\n", match.ElapsedTime.Seconds())
		fmt.Printf("Attack Count: %d\n", match.AttackCount)
		fmt.Println("Player Assists:")
		fmt.Printf("Success Rate: %.2f\n", match.SuccessRate)
		fmt.Println("---------------------------------")
	}
}
