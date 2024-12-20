package core

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/rustem-yam/esoft/internal/domain"
	"github.com/rustem-yam/esoft/internal/domain/errdomain"
)

func (c *Core) GetOffer(ctx context.Context, id int) (*domain.GetOfferResponse, error) {
	offer, err := c.db.GetOffer(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errdomain.OfferNotFoundError
		}
		return nil, errdomain.NewInternalError(err.Error())
	}
	return &domain.GetOfferResponse{Offer: *offer}, nil
}
