package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"noc-monitoring-bases/internal/database/connection"
	"noc-monitoring-bases/internal/database/models"
	"time"
)

type AuthPropertyRepositoryInterface interface {
	FindByCustomFilter(filter bson.M) ([]*models.AuthPropertyModel, error)
}

type AuthPropertyRepository struct {
	db *mongo.Collection
}

func NewAuthPropertyRepository() *AuthPropertyRepository {
	return &AuthPropertyRepository{
		db: connection.GetCollection(DbName, AuthProperty),
	}
}

func (r *AuthPropertyRepository) FindByCustomFilter(filter bson.M) ([]*models.AuthPropertyModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pipeline := []bson.M{
		{"$match": filter},
	}
	cursor, err := r.db.Aggregate(ctx, pipeline)

	if err != nil {
		return nil, fmt.Errorf("error finding documents: %v", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Printf("error closing cursor: %v", err)
		}
	}(cursor, ctx)

	var events []*models.AuthPropertyModel
	if err = cursor.All(ctx, &events); err != nil {
		return nil, fmt.Errorf("error en parseo cursor.All: %w", err)
	}
	return events, nil

}
