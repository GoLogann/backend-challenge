package usecase

import (
	"backend-challenge/internal/domain"
	"backend-challenge/internal/repository"
	"context"
)

type ProductUseCase struct {
	repo *repository.ProductRepository
}

func NewProductUseCase(repo *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, product *domain.Product) error {
	return uc.repo.Create(ctx, product)
}
