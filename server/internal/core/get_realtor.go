package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/rustem-yam/esoft/internal/domain"
	"github.com/rustem-yam/esoft/internal/domain/errdomain"
)

func (c *Core) GetRealtor(ctx context.Context, id int) (*domain.GetRealtorResponse, error) {
	response, err := c.db.GetRealtor(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.RealtorNotFoundError
		}
		return nil, errdomain.NewInternalError(err.Error())
	}
	return &domain.GetRealtorResponse{Realtor: *response}, nil
}
