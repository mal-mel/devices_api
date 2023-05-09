package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/mal-mel/devices_api/internal/api"
	"github.com/mal-mel/devices_api/internal/configs"
	"github.com/mal-mel/devices_api/internal/db"
	"github.com/mal-mel/devices_api/internal/logger"
	"github.com/mal-mel/devices_api/internal/server"
)

func main() {
	configPath := flag.String("c", "./config.yaml", "path to config")
	flag.Parse()

	ctx := context.Background()

	logger.ConfigureLogger()
	log := logrus.New()

	cfg := &configs.API{}
	err := configs.Read(*configPath, cfg)
	if err != nil {
		log.WithError(err).Fatal("can't read config")
	}

	log.WithField("config", cfg).Info("started with config")

	dbConn, err := db.ConnectPostgres(ctx, &cfg.DB)
	if err != nil {
		log.WithError(err).Error("can't connect database")
		return
	}

	serviceEnv := &api.Env{
		Log: log,
		DB:  &dbConn,
	}

	serverParam := api.ServerParam{
		APIEnv:            serviceEnv,
		TimeoutMiddleware: cfg.Timeout,
		ServerAddr:        cfg.ServeAddr,
		ReadinessPath:     cfg.Probes.Readiness,
		LivenessPath:      cfg.Probes.Liveness,
	}

	s := server.NewServer(log, &serverParam)
	s.Run()

	sgnl := make(chan os.Signal, 1)
	signal.Notify(sgnl,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	stop := <-sgnl
	s.Stop(ctx)
	log.WithField("signal", stop).Info("stopping")
}
