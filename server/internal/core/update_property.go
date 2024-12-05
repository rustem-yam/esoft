package core

import (
	"context"

	"github.com/rustem-yam/esoft/internal/domain"
	"github.com/rustem-yam/esoft/internal/domain/errdomain"
)

func (c *Core) UpdateProperty(ctx context.Context, req *domain.UpdatePropertyRequest) (*domain.GetPropertyResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	_, err = c.GetProperty(ctx, req.PropertyId)
	if err != nil {
		return nil, err
	}

	err = c.db.UpdateProperty(ctx, req)
	if err != nil {
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetProperty(ctx, req.PropertyId)
}
