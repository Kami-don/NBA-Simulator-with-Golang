package team

import (
	"context"

	"github.com/Furkan-Gulsen/NBA-Simulator-with-Golang/internal/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type RepositoryI interface {
	GetAll(tx context.Context) ([]*entities.Team, error)
	Get(tx context.Context, id int) (*entities.Team, error)
	Create(tx context.Context, t *entities.Team) error
	Update(tx context.Context, t *entities.Team) error
	Delete(tx context.Context, id int) error
}

type Repository struct {
	collection *mongo.Collection
}

func NewTeamRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}

func (r *Repository) GetAll(tx context.Context) ([]*entities.Team, error) {
	var teams []*entities.Team
	cursor, err := r.collection.Find(tx, nil)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(tx, &teams); err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *Repository) Get(tx context.Context, id int) (*entities.Team, error) {
	var team *entities.Team
	if err := r.collection.FindOne(tx, bson.M{"_id": id}).Decode(&team); err != nil {
		return nil, err
	}
	return team, nil
}

func (r *Repository) Create(tx context.Context, t *entities.Team) error {
	_, err := r.collection.InsertOne(tx, t)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Update(tx context.Context, t *entities.Team) error {
	_, err := r.collection.UpdateOne(tx, nil, bson.M{"$set": t})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(tx context.Context, id int) error {
	_, err := r.collection.DeleteOne(tx, nil)
	if err != nil {
		return err
	}
	return nil
}
