package entities

type Team struct {
	ID           int      `json:"id,omitempty" bson:"_id,omitempty"`
	Name         string   `json:"name,omitempty" bson:"name,omitempty"`
	Abbreviation string   `json:"abbreviation,omitempty" bson:"abbreviation,omitempty"`
	Players      []Player `json:"players,omitempty" bson:"-"`
}

// type TeamStats struct {
// 	// GameID      int `json:"game_id,omitempty" bson:"game_id,omitempty"`
// 	Score       int `json:"score,omitempty" bson:"score,omitempty"`
// 	AttackCount int `json:"attack_count,omitempty" bson:"attack_count,omitempty"`
// }

// type TeamGameInfo struct {
// 	// GameID      int           `json:"game_id,omitempty" bson:"game_id,omitempty"`
// 	Team        Team          `json:"team,omitempty" bson:"-"`
// 	PlayerStats []PlayerStats `json:"players,omitempty" bson:"-"`
// 	TeamStats   TeamStats     `json:"team_stats,omitempty" bson:"-"`
// }
