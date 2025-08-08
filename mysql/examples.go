package mysql

import (
	"context"
	"database/sql"
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
		VALUES (?, ?, ?);
	`
	res, err := db.ExecContext(ctx, q, first, last, email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetUsers(ctx context.Context, db *sql.DB, limit int) ([]User, error) {
	const q = `
		SELECT id, first_name, last_name, email
		FROM users
		ORDER BY id
		LIMIT ?;
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
	const q = `UPDATE users SET email = ? WHERE id = ?;`
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

	res, err := tx.ExecContext(ctx, `
		INSERT INTO users (first_name, last_name, email)
		VALUES (?, ?, ?)`, "Grace", "Hopper", "grace@example.com")
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `
		UPDATE users SET email = ? WHERE id = ?`,
		"grace.hopper@example.com", id); err != nil {
		return err
	}

	return tx.Commit()
}
