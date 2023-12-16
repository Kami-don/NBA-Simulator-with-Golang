package entities

type Player struct {
	ID     int           `json:"id" bson:"_id"`
	Name   string        `json:"name,omitempty" bson:"name,omitempty"`
	TeamID int           `json:"team_id,omitempty" bson:"team_id,omitempty"`
	Stats  []PlayerStats `json:"stats,omitempty" bson:"-"`
}

type PlayerStats struct {
	ID              int `json:"id" bson:"_id"`
	GameID          int `json:"game_id,omitempty" bson:"game_id,omitempty"`
	PlayerID        int `json:"player_id,omitempty" bson:"player_id,omitempty"`
	TwoPointScore   int `json:"two_point_score,omitempty" bson:"two_point_score,omitempty"`
	ThreePointScore int `json:"three_point_score,omitempty" bson:"three_point_score,omitempty"`
	Assist          int `json:"assist,omitempty" bson:"assist,omitempty"`
}
