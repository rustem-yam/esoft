package database

import (
	"context"
	"github.com/rustem-yam/esoft/internal/domain"
)

func (d *Database) GetClients(ctx context.Context) ([]domain.Client, error) {
	rows, err := d.Conn.QueryContext(ctx, `SELECT id, name, surname, patronymic, tel, email FROM clients`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []domain.Client

	for rows.Next() {
		var client domain.Client
		err := rows.Scan(
			&client.Id,
			&client.Name,
			&client.Surname,
			&client.Patronymic,
			&client.Telephone,
			&client.Email,
		)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}
