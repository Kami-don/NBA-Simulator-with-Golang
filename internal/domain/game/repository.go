package game

import (
	"context"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type RepositoryI interface {
	GetAll(tx context.Context) ([]*entities.Game, error)
	Get(tx context.Context, id string) (*entities.Game, error)
	Create(tx context.Context, g *entities.Game) error
	Update(tx context.Context, g *entities.Game) error
	Delete(tx context.Context, id string) error
}

type Repository struct {
	collection *mongo.Collection
}

func NewGameRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}

func (r *Repository) GetAll(tx context.Context) ([]*entities.Game, error) {
	var games []*entities.Game
	cursor, err := r.collection.Find(tx, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(tx, &games); err != nil {
		return nil, err
	}
	return games, nil
}

func (r *Repository) Get(tx context.Context, id string) (*entities.Game, error) {
	var game *entities.Game
	err := r.collection.FindOne(tx, nil).Decode(&game)
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
	_, err := r.collection.UpdateOne(tx, nil, bson.M{"$set": g})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(tx context.Context, id string) error {
	_, err := r.collection.DeleteOne(tx, nil)
	if err != nil {
		return err
	}
	return nil
}
