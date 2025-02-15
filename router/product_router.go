package router

import (
	"backend-challenge/internal/handler"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, productHandler *handler.ProductHandler) {
	product := r.Group("api/v1/products")
	{
		product.POST("", productHandler.CreateProduct)
	}
}
