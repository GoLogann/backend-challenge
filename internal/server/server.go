package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func StartServer(lc fx.Lifecycle, r *gin.Engine) {
	port := ":8080"
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logrus.Infof("API running on port %s", port)
				if err := r.Run(port); err != nil {
					logrus.Errorf("Error starting the server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logrus.Info("API has been stopped")
			return nil
		},
	})
}
