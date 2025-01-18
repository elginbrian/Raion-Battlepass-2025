package repository

import (
	"context"
	"fiber-starter/internal/domain"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

type authRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email, password_hash FROM users WHERE email = $1"

	row := r.db.QueryRow(ctx, query, email)

	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error scanning row: %w", err)
	}
    
	return &user, nil
}