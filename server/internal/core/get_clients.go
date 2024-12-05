package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/bouhartsev/infinity_realty/internal/domain"
	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
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
