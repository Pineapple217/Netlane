package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/Pineapple217/Netlane/pkg/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e      *echo.Echo
	config config.Server
}

func NewServer(conf config.Server) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = conf.Debug
	NewServer := &Server{
		e:      e,
		config: conf,
	}

	return NewServer
}

// Starts the server in a new routine
func (s *Server) Start() {
	slog.Info("Starting server")
	address := fmt.Sprintf("%s:%d", s.config.Bind, s.config.Port)
	go func() {
		if err := s.e.Start(address); err != nil && err != http.ErrServerClosed {
			slog.Error("Shutting down the server", "error", err.Error())
		}
	}()
	slog.Info("Server started", "address", address)
}

// Tries to the stops the server gracefully
func (s *Server) Stop() {
	slog.Info("Stopping server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
	}
}
