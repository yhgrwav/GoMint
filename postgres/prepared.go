package postgres

import (
	"context"
	"database/sql"
)

func PreparedInsertMany(ctx context.Context, db *sql.DB, us []User) error {
	const q = `
		INSERT INTO users (first_name, last_name, email)
		VALUES ($1, $2, $3)
		ON CONFLICT (email) DO NOTHING;
	`
	stmt, err := db.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, u := range us {
		if _, err := stmt.ExecContext(ctx, u.FirstName, u.LastName, u.Email); err != nil {
			return err
		}
	}
	return nil
}
