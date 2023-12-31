package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dragonator/evaluation-service/pkg/config"
)

// Logger is a contract to a logger implementation.
type Logger interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

// Service holds functionality for starting and stopping an HTTP server.
type Service struct {
	logger Logger
	server *http.Server
}

// New is a construction function for the HTTP server.
func New(config *config.Config, logger Logger, router http.Handler) (*Service, error) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.ServerPort),
		Handler: router,
	}

	return &Service{
		logger: logger,
		server: srv,
	}, nil
}

// Start starts the server.
func (s *Service) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatalf("ListenAndServe() failed: %v", err)
		}
	}()

	s.logger.Infof("Listening on port %s ...", s.server.Addr)
}

// Stop stops the server.
func (s *Service) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		s.logger.Fatalf("Shutdown() failed: %v", err)
	}

	s.logger.Info("HTTP server shut down gracefully.")
}
