package repository

import (
	"context"
	"errors"
	"fiber-starter/internal/domain"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, id string) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	UpdateUser(ctx context.Context, id string, user domain.User) (domain.User, error)
	DeleteUser(ctx context.Context, id string) error
	SearchUsers(ctx context.Context, query string) ([]domain.User, error)
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %w", err)
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	err := r.db.QueryRow(ctx, "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, errors.New("user not found")
		}
		return user, fmt.Errorf("error fetching user: %w", err)
	}
	return user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	commandTag, err := r.db.Exec(ctx, 
		"INSERT INTO users (name, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW())", 
		user.Name, user.Email, user.PasswordHash)
	if err != nil {
		return user, fmt.Errorf("error creating user: %w", err)
	}

	if commandTag.RowsAffected() > 0 {
		user.CreatedAt = time.Now()
		user.UpdatedAt = user.CreatedAt
		return user, nil
	}

	return user, errors.New("failed to create user")
}

func (r *userRepository) UpdateUser(ctx context.Context, id string, user domain.User) (domain.User, error) {
	commandTag, err := r.db.Exec(ctx, 
		"UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3", 
		user.Name, user.Email, id)
	if err != nil {
		return user, fmt.Errorf("error updating user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return user, errors.New("user not found")
	}

	user.ID = id
	user.UpdatedAt = time.Now()
	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	commandTag, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepository) SearchUsers(ctx context.Context, query string) ([]domain.User, error) {
	rows, err := r.db.Query(ctx, 
		"SELECT id, name, email, created_at, updated_at FROM users WHERE LOWER(name) LIKE LOWER($1) OR LOWER(email) LIKE LOWER($1)", 
		"%"+query+"%")	
    if err != nil {
        return nil, fmt.Errorf("error searching users: %w", err)
    }
    defer rows.Close()

    var users []domain.User
    for rows.Next() {
        var user domain.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
            return nil, fmt.Errorf("error scanning user row: %w", err)
        }
        users = append(users, user)
    }

    if len(users) == 0 {
        return nil, errors.New("no users found")
    }

    return users, nil
}