package playergame

import (
	"context"
	"fmt"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type RepositoryI interface {
	Get(ctx context.Context, id int) (*entities.PlayerStats, error)
	GetAll(ctx context.Context) ([]*entities.PlayerStats, error)
	Create(ctx context.Context, p *entities.PlayerStats) error
	Update(ctx context.Context, p *entities.PlayerStats) error
	Delete(ctx context.Context, id int) error
	GetAssistLeader(ctx context.Context) (*entities.PlayerStats, error)
	GetPlayersByGameID(ctx context.Context, teamID int) ([]*entities.PlayerStats, error)
}

type Repository struct {
	collection *mongo.Collection
}

func NewPlayerGameRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}

func (r *Repository) Get(ctx context.Context, id int) (*entities.PlayerStats, error) {
	var player *entities.PlayerStats
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&player)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]*entities.PlayerStats, error) {
	var players []*entities.PlayerStats
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &players); err != nil {
		return nil, err
	}
	return players, nil
}

func (r *Repository) Create(ctx context.Context, p *entities.PlayerStats) error {
	_, err := r.collection.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, p *entities.PlayerStats) error {
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

func (r *Repository) GetAssistLeader(ctx context.Context) (*entities.PlayerStats, error) {
	leader := entities.PlayerStats{}
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var p entities.PlayerStats
		err := cursor.Decode(&p)
		if err != nil {
			fmt.Println(err)
		}

		if p.Assist > leader.Assist {
			leader = p
		}
	}

	cursor.Close(ctx)

	return &leader, nil
}

func (r *Repository) GetPlayersByGameID(ctx context.Context, gameID int) ([]*entities.PlayerStats, error) {
	var players []*entities.PlayerStats
	cursor, err := r.collection.Find(ctx, bson.M{"game_id": gameID})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &players); err != nil {
		return nil, err
	}
	return players, nil
}
