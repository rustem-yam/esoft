package database

import (
	"context"
	"database/sql"
	"go.uber.org/zap"
)

func (d *Database) CheckClientEmail(ctx context.Context, email string) error {
	var userId int
	d.Logger.Info("Executing CheckClientEmail", zap.String("email", email))

	row := d.Conn.QueryRowContext(ctx, "SELECT id FROM clients WHERE LOWER(email) = LOWER(?)", email)
	err := row.Scan(&userId)

	if err == sql.ErrNoRows {
		d.Logger.Info("No client found with this email", zap.String("email", email))
		return err
	}

	if err != nil {
		d.Logger.Error("Error executing CheckClientEmail", zap.Error(err))
		return err
	}

	d.Logger.Info("Client found with email", zap.Int("userId", userId))
	return nil
}
