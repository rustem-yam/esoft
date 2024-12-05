package database

import (
	"context"
	"go.uber.org/zap"
	"github.com/rustem-yam/esoft/internal/domain"
)

func (d *Database) CreateClient(ctx context.Context, req *domain.CreateClientRequest) (int, error) {
	d.Logger.Info("Creating a new client", zap.String("name", req.Name), zap.String("surname", req.Surname))

	res, err := d.Conn.ExecContext(ctx, `insert into clients(name, surname, patronymic, tel, email)
							 values(?, ?, ?, ?, ?)`,
		req.Name,
		req.Surname,
		req.Patronymic,
		req.Telephone,
		req.Email,
	)
	if err != nil {
		d.Logger.Error("Failed to create client", zap.Error(err))
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		d.Logger.Error("Failed to retrieve last insert ID", zap.Error(err))
		return 0, err
	}

	d.Logger.Info("Client created successfully", zap.Int64("clientID", id))

	return int(id), nil
}
