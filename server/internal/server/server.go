package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rustem-yam/esoft/internal/config"
	"github.com/rustem-yam/esoft/internal/core"
	"github.com/rustem-yam/esoft/internal/domain/errdomain"
	"github.com/rustem-yam/esoft/internal/persistence/database"

	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	cfg    *config.Config
	core   *core.Core
	db     *database.Database
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func New(l *zap.Logger, c *config.Config) (*Server, error) {
	s := &Server{
		logger: l,
		cfg:    c,
	}

	db, err := database.NewMySQLConnection(s.cfg.DatabaseConnection)
	if err != nil {
		return nil, err
	}

	l.Info("Connected to database", zap.String("conn", c.DatabaseConnection))
	s.core = core.NewCore(l, db, c)

	return s, nil
}

func (s *Server) Run() error {
	router := s.initRoutes()

	handler := enableCORS(router)

	httpServer := &http.Server{
		Addr:    s.cfg.AppPort,
		Handler: handler,
	}

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("failed to listen and serve", zap.Error(err), zap.String("address", httpServer.Addr))
			quit <- os.Interrupt
		}
	}()

	s.logger.Info("Running the Server", zap.String("address", httpServer.Addr))

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return func() error {
		s.logger.Info("shutting down the Server...")

		err := httpServer.Shutdown(ctx)

		if err != nil {
			s.logger.Error(err.Error())
			return err
		}

		s.logger.Info("Server had shut down successfully")

		return nil
	}()
}

func (s *Server) ParseError(err error, w http.ResponseWriter) {
	var e errdomain.AppError

	if errors.As(err, &e) {
		s.Json(e, e.HttpCode, w)
	} else {
		s.Json(&errdomain.AppError{Message: err.Error()}, http.StatusInternalServerError, w)
	}
}
