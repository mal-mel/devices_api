package server

import (
	"context"
	"net/http"

	"github.com/mal-mel/devices_api/internal/api"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger

	httpServer *http.Server
}

func NewServer(log *logrus.Logger, serverParam *api.ServerParam) *Server {
	httpServer, err := api.NewServer(serverParam)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		logger:     log,
		httpServer: httpServer,
	}
}

func (s *Server) Run() {
	go func() {
		s.logger.Info("script manager service started on port", s.httpServer.Addr)

		if err := s.httpServer.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				s.logger.Info("graceful shutdown")
			} else {
				s.logger.WithError(err).Fatal("script manager service")
			}
		}
	}()
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		s.logger.WithError(err).Error("error on http server shutdown")
	}
}
