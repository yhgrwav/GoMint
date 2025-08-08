package postgres

import (
	"context"
	"database/sql"
)

func GetUserByEmail(ctx context.Context, db *sql.DB, email string) (User, error) {
	const q = `
		SELECT id, first_name, last_name, email
		FROM users
		WHERE email = $1;
	`
	var u User
	err := db.QueryRowContext(ctx, q, email).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
