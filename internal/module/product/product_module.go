package product

import (
	"backend-challenge/internal/handler"
	"backend-challenge/internal/repository"
	"backend-challenge/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewProductRepository),
	fx.Provide(NewProductUseCase),
	fx.Provide(NewProductHandler),
)

func NewProductRepository(db *mongo.Database) *repository.ProductRepository {
	return repository.NewProductRepository(db)
}

func NewProductUseCase(productRepo *repository.ProductRepository) *usecase.ProductUseCase {
	return usecase.NewProductUseCase(productRepo)
}

func NewProductHandler(pu *usecase.ProductUseCase) *handler.ProductHandler {
	return handler.NewProductHandler(pu)
}
