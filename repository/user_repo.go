package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	db *pgx.Conn
}

func ProvideUserRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) CheckUserExists(ctx context.Context, account, email string) (exists bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE account = $1 OR email = $2)`
	err = repo.db.QueryRow(ctx, query, account, email).Scan(&exists)
	return exists, err
}

func (repo *UserRepo) CreateUser(ctx context.Context, account, password, email string) (suc bool, err error) {
	query := `
		INSERT INTO users (account, pwd, email)
		VALUES ($1, $2, $3)
	`
	_, err = repo.db.Exec(ctx, query, account, password, email)
	if err != nil {
		return false, err
	}
	return true, nil
}
