package core

import (
	"context"

	"github.com/rustem-yam/esoft/internal/domain"
	"github.com/rustem-yam/esoft/internal/domain/errdomain"
)

func (c *Core) CreateOffer(ctx context.Context, req *domain.CreateOfferRequest) (*domain.GetOfferResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	id, err := c.db.CreateOffer(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetOffer(ctx, id)
}
