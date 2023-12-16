package entities

import "time"

type Game struct {
	ID          int          `json:"id" bson:"_id"`
	Away        TeamGameInfo `json:"away,omitempty" bson:"away,omitempty"` // one-to-one relationship
	Home        TeamGameInfo `json:"home,omitempty" bson:"home,omitempty"` // one-to-one relationship
	StartTime   time.Time    `json:"start_time,omitempty" bson:"start_time,omitempty"`
	ElapsedTime int          `json:"elapsed_time,omitempty" bson:"elapsed_time,omitempty"`
	IsFinished  bool         `json:"is_finished,omitempty" bson:"is_finished,omitempty"`
}

type TeamGameInfo struct {
	TeamID      int           `json:"team_id,omitempty" bson:"team_id,omitempty"`
	PlayerStats []PlayerStats `json:"players,omitempty" bson:"-"`
	Score       int           `json:"score,omitempty" bson:"score,omitempty"`
	AttackCount int           `json:"attack_count,omitempty" bson:"attack_count,omitempty"`
	SuccessRate float64       `json:"success_rate,omitempty" bson:"success_rate,omitempty"`
}

type Scoreboard struct {
	HomeScore       int     `json:"home_score"`
	AwayScore       int     `json:"away_score"`
	HomeTeam        string  `json:"home_team"`
	AwayTeam        string  `json:"away_team"`
	HomeSuccessRate float64 `json:"home_success_rate"`
	AwaySuccessRate float64 `json:"away_success_rate"`
}
