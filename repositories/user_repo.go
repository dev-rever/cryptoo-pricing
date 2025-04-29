package repositories

import (
	"context"

	"github.com/dev-rever/cryptoo-pricing/model"
	"github.com/jackc/pgx/v5"
)

type User struct {
	db *pgx.Conn
}

func ProvideUserRepo(db *pgx.Conn) *User {
	return &User{db: db}
}

func (repo *User) CheckUserExists(ctx context.Context, account, email string) (exists bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE account = $1 OR email = $2)`
	err = repo.db.QueryRow(ctx, query, account, email).Scan(&exists)
	return exists, err
}

func (repo *User) QueryUserIDByAccount(ctx context.Context, account string) (uid uint, err error) {
	query := `SELECT id FROM users WHERE account = $1`
	err = repo.db.QueryRow(ctx, query, account).Scan(&uid)
	return uid, err
}

func (repo *User) QueryUserPwdByAccount(ctx context.Context, account string) (pwd string, err error) {
	query := `SELECT pwd FROM users WHERE account = $1`
	err = repo.db.QueryRow(ctx, query, account).Scan(&pwd)
	return pwd, err
}

func (repo *User) InsertUser(ctx context.Context, account, password, email string) (uid uint, err error) {
	query := `INSERT INTO users (account, pwd, email) VALUES ($1, $2, $3) RETURNING id`
	err = repo.db.QueryRow(ctx, query, account, password, email).Scan(&uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

func (repo *User) QueryUserByID(ctx context.Context, uid uint) (user model.UserInfo, err error) {
	query := `SELECT * FROM users WHERE id = $1`
	err = repo.db.QueryRow(ctx, query, uid).Scan(&user.ID, &user.Account, &user.Password, &user.Email)
	return user, err
}
