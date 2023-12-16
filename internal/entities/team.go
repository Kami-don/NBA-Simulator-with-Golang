package entities

type Team struct {
	Name    string   `json:"name,omitempty" bson:"name,omitempty"`
	Players []Player `json:"players,omitempty" bson:"players,omitempty"`
}

type TeamStats struct {
	GameID string `json:"game_id,omitempty" bson:"game_id,omitempty"`
	Score  int    `json:"score,omitempty" bson:"score,omitempty"`
}

type TeamGameInfo struct {
	ID        string           `json:"id,omitempty" bson:"id,omitempty"`
	Team      Team             `json:"team,omitempty" bson:"team,omitempty"`
	Players   []PlayerGameInfo `json:"players,omitempty" bson:"players,omitempty"`
	TeamStats TeamStats        `json:"team_stats,omitempty" bson:"team_stats,omitempty"`
}
