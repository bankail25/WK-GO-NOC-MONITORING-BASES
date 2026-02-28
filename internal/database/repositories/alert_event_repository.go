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

type AlertEventRepositoryInterface interface {
	Create(alertEvent *models.AlertEventModel) (*mongo.InsertOneResult, error)
	FindByCustomFilter(filter bson.M) ([]*models.AlertEventModel, error)
}

type AlertEventRepository struct {
	db *mongo.Collection
}

func NewAlertEventRepository() *AlertEventRepository {
	return &AlertEventRepository{
		db: connection.GetCollection(DbName, AlertEvents),
	}
}

func (r *AlertEventRepository) Create(alertEvent *models.AlertEventModel) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, err := r.db.InsertOne(ctx, alertEvent)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *AlertEventRepository) FindByCustomFilter(filter bson.M) ([]*models.AlertEventModel, error) {
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

	var events []*models.AlertEventModel
	if err = cursor.All(ctx, &events); err != nil {
		return nil, fmt.Errorf("error en parseo cursor.All: %w", err)
	}
	return events, nil

}
