package core

import (
	"context"

	"github.com/rustem-yam/esoft/internal/domain/errdomain"
)

func (c *Core) DeleteOffer(ctx context.Context, id int) error {
	_, err := c.GetOffer(ctx, id)
	if err != nil {
		return err
	}

	err = c.db.DeleteOffer(ctx, id)
	if err != nil {
		return errdomain.NewInternalError(err.Error())
	}

	return nil
}
