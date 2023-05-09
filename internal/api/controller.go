package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sirupsen/logrus"

	"github.com/mal-mel/devices_api/internal/db"
)

type Env struct {
	Log *logrus.Logger
	DB  db.Database
}

type ServerParam struct {
	APIEnv *Env

	TimeoutMiddleware time.Duration
	ServerAddr        string

	LivenessPath  string
	ReadinessPath string
}

func NewServer(param *ServerParam) (*http.Server, error) {
	mux := chi.NewRouter()

	mux.Use(middleware.NoCache)

	mux.Use(LoggerMiddleware(logrus.New()))

	mux.Use(middleware.Timeout(param.TimeoutMiddleware))

	mux.Get(param.LivenessPath, param.APIEnv.Alive)
	mux.Get(param.ReadinessPath, param.APIEnv.Ready)

	const (
		baseURL = "/api/v1"
	)

	mux.Mount(baseURL+"/pprof", middleware.Profiler())

	mux.With(middleware.SetHeader("Content-Type", "application/json")).
		Route(baseURL, func(r chi.Router) {
			HandlerFromMux(param.APIEnv, r)
		})

	return &http.Server{
		Addr:    param.ServerAddr,
		Handler: mux,
	}, nil
}
