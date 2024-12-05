package core

import (
	"context"

	"github.com/rustem-yam/esoft/internal/domain"
	"github.com/rustem-yam/esoft/internal/domain/errdomain"
)

func (c *Core) CreateRealtor(ctx context.Context, req *domain.CreateRealtorRequest) (*domain.GetRealtorResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	id, err := c.db.CreateRealtor(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetRealtor(ctx, id)
}
