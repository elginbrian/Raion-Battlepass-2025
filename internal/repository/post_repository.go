package repository

import (
	"context"
	"fiber-starter/internal/domain"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostRepository interface {
	FetchAllPosts(ctx context.Context) ([]domain.Post, error)
	FetchPostByID(ctx context.Context, postID string) (*domain.Post, error)
	FetchPostsByUserID(ctx context.Context, userID string) ([]domain.Post, error)
	CreatePost(ctx context.Context, post domain.Post) (*domain.Post, error)
	UpdatePost(ctx context.Context, postID string, post domain.Post) (*domain.Post, error)
	DeletePost(ctx context.Context, postID string) error
	SearchPosts(ctx context.Context, query string) ([]domain.Post, error)
}

type postRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) FetchAllPosts(ctx context.Context) ([]domain.Post, error) {
	query := "SELECT id, user_id, caption, image_url, created_at, updated_at FROM posts"
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching posts: %w", err)
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Caption, &post.ImageURL, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning post row: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}
	return posts, nil
}

func (r *postRepository) FetchPostByID(ctx context.Context, postID string) (*domain.Post, error) {
	query := "SELECT id, user_id, caption, image_url, created_at, updated_at FROM posts WHERE id = $1"
	row := r.db.QueryRow(ctx, query, postID)

	var post domain.Post
	if err := row.Scan(&post.ID, &post.UserID, &post.Caption, &post.ImageURL, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil 
		}
		return nil, fmt.Errorf("error fetching post by ID: %w", err)
	}
	return &post, nil
}

func (r *postRepository) FetchPostsByUserID(ctx context.Context, userID string) ([]domain.Post, error) {
	query := "SELECT id, user_id, caption, image_url, created_at, updated_at FROM posts WHERE user_id = $1"
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching posts for user %s: %w", userID, err)
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Caption, &post.ImageURL, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning post row: %w", err)
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}
	return posts, nil
}

func (r *postRepository) CreatePost(ctx context.Context, post domain.Post) (*domain.Post, error) {
	query := "INSERT INTO posts (user_id, caption, image_url) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	err := r.db.QueryRow(ctx, query, post.UserID, post.Caption, post.ImageURL).Scan(
		&post.ID, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating post: %w", err)
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = post.CreatedAt
	return &post, nil
}

func (r *postRepository) UpdatePost(ctx context.Context, postID string, post domain.Post) (*domain.Post, error) {
	query := "UPDATE posts SET caption = $1, image_url = $2, updated_at = NOW() WHERE id = $3 RETURNING id, user_id, caption, image_url, created_at, updated_at"
	err := r.db.QueryRow(ctx, query, post.Caption, post.ImageURL, postID).Scan(
		&post.ID, &post.UserID, &post.Caption, &post.ImageURL, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error updating post: %w", err)
	}
	return &post, nil
}

func (r *postRepository) DeletePost(ctx context.Context, postID string) error {
	query := "DELETE FROM posts WHERE id = $1"
	result, err := r.db.Exec(ctx, query, postID)
	if err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *postRepository) SearchPosts(ctx context.Context, query string) ([]domain.Post, error) {
	queryText := `
		SELECT id, user_id, caption, image_url, created_at, updated_at 
		FROM posts 
		WHERE LOWER(caption) LIKE LOWER($1) OR LOWER(image_url) LIKE LOWER($1)
	`
	rows, err := r.db.Query(ctx, queryText, "%"+query+"%")
	if err != nil {
		return nil, fmt.Errorf("error searching posts: %w", err)
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var post domain.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Caption, &post.ImageURL, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning post row: %w", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	if len(posts) == 0 {
		return nil, fmt.Errorf("no posts found matching query: %s", query)
	}

	return posts, nil
}