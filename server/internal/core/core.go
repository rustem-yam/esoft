package core

import (
	"database/sql"

	"go.uber.org/zap"

	"github.com/rustem-yam/esoft/internal/config"
	"github.com/rustem-yam/esoft/internal/persistence/database"
)

type Core struct {
	logger *zap.Logger
	db     *database.Database
	cfg    *config.Config
}

func NewCore(l *zap.Logger, db *sql.DB, c *config.Config) *Core {
	return &Core{
		logger: l,
		db:     &database.Database{Conn: db},
		cfg:    c,
	}
}
