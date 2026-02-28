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

type AlertRulesRepositoryInterface interface {
	FindAll(string) ([]*models.AlertRuleModel, error)
	FindByActive() ([]*models.AlertRuleModel, error)
	FindByCustomFilter(bson.M) ([]*models.AlertRuleWithContacts, error)
}

type AlertRuleRepository struct {
	db *mongo.Collection
}

func NewQueryRulesRepository() *AlertRuleRepository {
	return &AlertRuleRepository{
		db: connection.GetCollection(DbName, AlertRule),
	}
}

func (r *AlertRuleRepository) FindAll(db string) ([]*models.AlertRuleModel, error) {
	return nil, nil
}

func (r *AlertRuleRepository) FindByActive() ([]*models.AlertRuleModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"is_active": true}

	cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %v", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Printf("error closing cursor: %v", err)
		}
	}(cursor, ctx)

	var rules []*models.AlertRuleModel
	if err = cursor.All(ctx, &rules); err != nil {
		return nil, fmt.Errorf("error en parseo cursor.All: %w", err)
	}

	return rules, nil
}

func (r *AlertRuleRepository) FindByCustomFilter(filter bson.M) ([]*models.AlertRuleWithContacts, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pipeline := []bson.M{
		{"$match": filter},
	}
	cursor, err := r.db.Aggregate(ctx, pipeline)

	//cursor, err := r.db.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error finding documents: %v", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Printf("error closing cursor: %v", err)
		}
	}(cursor, ctx)

	var rules []*models.AlertRuleWithContacts
	if err = cursor.All(ctx, &rules); err != nil {
		return nil, fmt.Errorf("error en parseo cursor.All: %w", err)
	}

	return rules, nil

}
