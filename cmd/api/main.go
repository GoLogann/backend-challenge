package main

import (
	"backend-challenge/internal/config"
	"backend-challenge/internal/module/product"
	"backend-challenge/internal/repository"
	"backend-challenge/internal/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

func main() {
	cfg := config.LoadConfig()

	app := fx.New(
		fx.Provide(func() (*repository.MongoDB, error) {
			return repository.NewMongoDB(cfg.MongoURI, cfg.Database)
		}),
		fx.Provide(func(m *repository.MongoDB) *mongo.Database {
			return m.Database
		}),

		server.RouterModule(),
		product.Module,
		fx.Invoke(server.RegisterRoutes),
	)

	app.Run()
}
