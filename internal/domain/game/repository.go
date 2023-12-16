package game

import (
	"context"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type RepositoryI interface {
	GetAll(ctx context.Context) ([]*entities.Game, error)
	Get(ctx context.Context, id int) (*entities.Game, error)
	Create(ctx context.Context, g *entities.Game) error
	Update(ctx context.Context, g *entities.Game) error
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	collection *mongo.Collection
}

func NewGameRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}

func (r *Repository) GetAll(tx context.Context) ([]*entities.Game, error) {
	var games []*entities.Game
	cursor, err := r.collection.Find(tx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(tx, &games)
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (r *Repository) Get(tx context.Context, id int) (*entities.Game, error) {
	var game *entities.Game
	err := r.collection.FindOne(tx, bson.M{"_id": id}).Decode(&game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (r *Repository) Create(tx context.Context, g *entities.Game) error {
	_, err := r.collection.InsertOne(tx, g)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(tx context.Context, g *entities.Game) error {
	_, err := r.collection.UpdateOne(tx, bson.M{"_id": g.ID}, bson.M{"$set": g})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(tx context.Context, id int) error {
	_, err := r.collection.DeleteOne(tx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
