package postgres

import (
	"context"
	"database/sql"
	"errors"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
}

func InsertUser(ctx context.Context, db *sql.DB, first, last, email string) (int64, error) {
	const q = `
		INSERT INTO users (first_name, last_name, email)
		VALUES ($1, $2, $3)
		RETURNING id;
	`
	var id int64
	if err := db.QueryRowContext(ctx, q, first, last, email).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func GetUsers(ctx context.Context, db *sql.DB, limit int) ([]User, error) {
	const q = `
		SELECT id, first_name, last_name, email
		FROM users
		ORDER BY id
		LIMIT $1;
	`
	rows, err := db.QueryContext(ctx, q, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}

func UpdateEmail(ctx context.Context, db *sql.DB, id int64, email string) (int64, error) {
	const q = `UPDATE users SET email = $1 WHERE id = $2;`
	res, err := db.ExecContext(ctx, q, email, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func TxExample(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	var newID int64
	if err := tx.QueryRowContext(ctx, `
		INSERT INTO users (first_name, last_name, email)
		VALUES ($1,$2,$3) RETURNING id;`,
		"Grace", "Hopper", "grace@example.com",
	).Scan(&newID); err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `
		UPDATE users SET email = $1 WHERE id = $2;`,
		"grace.hopper@example.com", newID,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func SafeGetUser(ctx context.Context, tx *sql.Tx, id int64) (User, error) {
	const q = `
		SELECT id, first_name, last_name, email
		FROM users
		WHERE id = $1
		FOR UPDATE;
	`
	var u User
	err := tx.QueryRowContext(ctx, q, id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, err
	}
	return u, err
}
