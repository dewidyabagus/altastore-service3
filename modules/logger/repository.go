package logger

import (
	"context"

	"AltaStore/business/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	DB *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{db}
}

func (r *Repository) AddLoggerActivity(data *logger.LoggerData) error {

	_, err := r.DB.Collection("logs").InsertOne(context.Background(), data)

	return err
}
