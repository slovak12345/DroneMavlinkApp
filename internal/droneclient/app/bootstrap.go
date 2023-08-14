package app

import (
	"log"
	"time"

	"github.com/slovak12345/DroneMavlinkApp/internal/droneclient/handlers"
	"github.com/slovak12345/DroneMavlinkApp/internal/droneclient/infrastructure/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func timeFormat(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		loc, _ = time.LoadLocation("Local")
	}
	enc.AppendString(t.In(loc).Format("2006-01-02 15:04:05.000"))
}

func (a *App) Bootstrap() error {
	conf := zap.NewProductionConfig()
	conf.EncoderConfig.EncodeTime = timeFormat

	logger, err := conf.Build()
	if err != nil {
		log.Fatal("failed to initialize logger")
	}
	a.log = logger

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	a.cfg = cfg

	a.log.Info("DroneWebService is bootstrapping")

	a.handler = handlers.NewHandler(a.log)

	a.log.Info("DroneWebService bootstrapped")
	return nil
}
