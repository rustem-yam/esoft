package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/rustem-yam/esoft/internal/domain"
	"github.com/rustem-yam/esoft/internal/domain/errdomain"
)

func (c *Core) GetClients(ctx context.Context) (*domain.GetClientsResponse, error) {
	clients, err := c.db.GetClients(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.ClientsNotFoundError
		}
		return nil, errdomain.NewInternalError(err.Error())
	}

	return &domain.GetClientsResponse{Clients: clients}, nil
}
