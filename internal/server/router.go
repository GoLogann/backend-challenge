package server

import (
	"backend-challenge/internal/handler"
	"backend-challenge/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func RegisterRoutes(
	r *gin.Engine,
	productHandler *handler.ProductHandler,
) {
	router.ProductRoutes(r, productHandler)
}

func RouterModule() fx.Option {
	return fx.Options(
		fx.Provide(gin.Default()),
		fx.Invoke(StartServer),
	)
}
