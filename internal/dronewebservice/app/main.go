package app

import (
	"log"

	"github.com/slovak12345/DroneMavlinkApp/internal/dronewebservice/handlers"
	"github.com/slovak12345/DroneMavlinkApp/internal/dronewebservice/infrastructure/config"
	"github.com/slovak12345/DroneMavlinkApp/pkg/graceful"
	"go.uber.org/zap"
)

type App struct {
	log     *zap.Logger
	cfg     *config.Config
	handler *handlers.Handler
}

func Start() {
	app := new(App)
	if err := app.Bootstrap(); err != nil {
		log.Fatal(err)
	}

	go func() {
		app.Run()
	}()

	err := graceful.WaitShutdown()

	if err != nil {
		app.log.Fatal("DroneWebService is dead")
	} else {
		app.log.Info("DroneWebService gracefully stopped")
	}
}
