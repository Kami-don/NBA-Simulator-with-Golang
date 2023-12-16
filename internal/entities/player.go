package entities

type Player struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Team string `json:"team,omitempty" bson:"team,omitempty"`
}

type PlayerStats struct {
	TwoPointScore   int `json:"two_point_score,omitempty" bson:"two_point_score,omitempty"`
	ThreePointScore int `json:"three_point_score,omitempty" bson:"three_point_score,omitempty"`
	Assist          int `json:"assist,omitempty" bson:"assist,omitempty"`
}

type PlayerGameInfo struct {
	ID     string      `json:"id" bson:"_id"`
	GameID string      `json:"game_id,omitempty" bson:"game_id,omitempty"`
	Player Player      `json:"player,omitempty" bson:"player,omitempty"`
	Stats  PlayerStats `json:"stats,omitempty" bson:"stats,omitempty"`
}
