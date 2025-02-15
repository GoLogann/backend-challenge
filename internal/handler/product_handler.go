package handler

import (
	"backend-challenge/internal/domain"
	"backend-challenge/internal/usecase"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	useCase *usecase.ProductUseCase
}

func NewProductHandler(useCase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{useCase: useCase}
}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ph.useCase.CreateProduct(context.Background(), &product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
