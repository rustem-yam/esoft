package core

import (
	"context"
	"errors"

	"github.com/rustem-yam/esoft/internal/domain/errdomain"
	"github.com/rustem-yam/esoft/internal/persistence/database"
)

func (c *Core) DeleteProperty(ctx context.Context, id int) error {
	_, err := c.GetProperty(ctx, id)
	if err != nil {
		return err
	}

	err = c.db.CheckPropertyDependency(ctx, id)
	if err != nil {
		if errors.Is(err, database.ErrPropertyHasOffers) {
			return errdomain.PropertyHasOffersError
		}
		return errdomain.NewInternalError(err.Error())
	}

	err = c.db.DeleteProperty(ctx, id)
	if err != nil {
		return errdomain.NewInternalError(err.Error())
	}

	return nil
}
