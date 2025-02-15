package repository

import (
	"backend-challenge/internal/domain"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{collection: db.Collection("products")}
}

func (r *ProductRepository) Create(ctx context.Context, product *domain.Product) error {
	_, err := r.collection.InsertOne(ctx, product)
	return err
}

func (r *ProductRepository) FindByOwner(ctx context.Context, ownerID string) ([]domain.Product, error) {
	var products []domain.Product
	cursor, err := r.collection.Find(ctx, bson.M{"owner_id": ownerID})
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) error {
		err = cursor.Close(ctx)
		if err != nil {
			return errors.New("failed to close cursor")
		}
		return nil
	}(cursor, ctx)
	for cursor.Next(ctx) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
