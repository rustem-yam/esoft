package core

import (
	"context"
	"errors"

	"database/sql"

	"github.com/rustem-yam/esoft/internal/domain"
	"github.com/rustem-yam/esoft/internal/domain/errdomain"
)

func (c *Core) CreateClient(ctx context.Context, req *domain.CreateClientRequest) (*domain.GetClientResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	c.logger.Info("creating clientttt")

	if req.Email != nil {
		err := c.db.CheckClientEmail(ctx, *req.Email)
		c.logger.Info("Checking client email")
		
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, errdomain.NewInternalError(err.Error())
		}
		
		if err == nil {
			return nil, errdomain.NewUserError("Email already taken.", "email")
		}
	}
	

	if req.Telephone != nil {
		err := c.db.CheckClientTelephone(ctx, *req.Telephone)
		c.logger.Info("error clienttt2")
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, errdomain.NewInternalError(err.Error())
		}
		if err == nil {
			return nil, errdomain.NewUserError("Telephone already taken.", "telephone")
		}
	}

	id, err := c.db.CreateClient(ctx, req)
	if err != nil {
		c.logger.Info("error clienttt3")
		return nil, errdomain.NewInternalError(err.Error())
	}

	return c.GetClient(ctx, id)
}
