package service

import (
	"context"
	"fiber-starter/internal/domain"
	"fiber-starter/internal/repository"
)

type PostService interface {
	FetchAllPosts() ([]domain.Post, error)
	FetchPostByID(id string) (domain.Post, error)
	FetchPostsByUserID(userID string) ([]domain.Post, error)
	CreatePost(post domain.Post) (domain.Post, error)
	UpdatePost(id string, post domain.Post) (domain.Post, error)
	DeletePost(id string) error
	SearchPosts(query string) ([]domain.Post, error)
}

type postService struct {
	postRepo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{postRepo: repo}
}

func (s *postService) FetchAllPosts() ([]domain.Post, error) {
	ctx := context.Background()
	return s.postRepo.FetchAllPosts(ctx)
}

func (s *postService) FetchPostByID(id string) (domain.Post, error) {
	ctx := context.Background()
	post, err := s.postRepo.FetchPostByID(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}
	if post == nil {
		return domain.Post{}, domain.ErrNotFound
	}
	return *post, nil
}

func (s *postService) FetchPostsByUserID(userID string) ([]domain.Post, error) {
	ctx := context.Background()
	posts, err := s.postRepo.FetchPostsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return nil, domain.ErrNotFound 
	}
	return posts, nil
}

func (s *postService) CreatePost(post domain.Post) (domain.Post, error) {
	ctx := context.Background()
	createdPost, err := s.postRepo.CreatePost(ctx, post)
	if err != nil {
		return domain.Post{}, err
	}
	if createdPost == nil {
		return domain.Post{}, domain.ErrNotFound
	}
	return *createdPost, nil
}

func (s *postService) UpdatePost(id string, post domain.Post) (domain.Post, error) {
	ctx := context.Background()
	updatedPost, err := s.postRepo.UpdatePost(ctx, id, post)
	if err != nil {
		return domain.Post{}, err
	}
	if updatedPost == nil {
		return domain.Post{}, domain.ErrNotFound
	}
	return *updatedPost, nil
}

func (s *postService) DeletePost(id string) error {
	ctx := context.Background()
	err := s.postRepo.DeletePost(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *postService) SearchPosts(query string) ([]domain.Post, error) {
	ctx := context.Background()
	posts, err := s.postRepo.SearchPosts(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return nil, domain.ErrNotFound 
	}
	return posts, nil
}