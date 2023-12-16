package player

import (
	"context"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type RepositoryI interface {
	Get(ctx context.Context, id int) (*entities.Player, error)
	GetAll(ctx context.Context) ([]*entities.Player, error)
	Create(ctx context.Context, p *entities.Player) error
	Update(ctx context.Context, p *entities.Player) error
	Delete(ctx context.Context, id int) error
	GetPlayersByTeamID(ctx context.Context, teamID int) ([]*entities.Player, error)
}

type Repository struct {
	collection *mongo.Collection
}

func NewPlayerRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}

func (r *Repository) Get(ctx context.Context, id int) (*entities.Player, error) {
	var player *entities.Player
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]*entities.Player, error) {
	var players []*entities.Player
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &players); err != nil {
		return nil, err
	}
	return players, nil
}

func (r *Repository) Create(ctx context.Context, p *entities.Player) error {
	_, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, p *entities.Player) error {
	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": p.ID}, bson.M{"$set": p})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetPlayersByTeamID(ctx context.Context, teamID int) ([]*entities.Player, error) {
	var players []*entities.Player
	cursor, err := r.collection.Find(ctx, bson.M{"team_id": teamID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &players); err != nil {
		return nil, err
	}
	return players, nil
}
