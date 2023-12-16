package entities

type Player struct {
	ID   int64  `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Team string `json:"team" bson:"team"`
}
