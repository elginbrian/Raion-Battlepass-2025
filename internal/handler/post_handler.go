package handler

import (
	"fiber-starter/internal/domain"
	"fiber-starter/internal/service"
	"fiber-starter/pkg/request"
	"fiber-starter/pkg/response"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	postService service.PostService
	authService service.AuthService
}

func NewPostHandler(postService service.PostService, authService service.AuthService) *PostHandler {
	return &PostHandler{
        postService: postService,
        authService: authService,
    }
}


// GetAllPosts godoc
// @Summary Get all posts
// @Description Get a list of all posts, along with details like the user who created them, the caption, image URL, and timestamps.
// @Tags posts
// @Produce json
// @Success 200 {object} response.GetAllPostsResponse "Successful fetch posts response"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/posts [get]
func (h *PostHandler) GetAllPosts(c *fiber.Ctx) error {
	posts, err := h.postService.FetchAllPosts()
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	var postResponse []domain.PostResponse
	for _, post := range posts {
		postResponse = append(postResponse, domain.PostResponse{
			ID:        post.ID,
			UserID:    post.UserID,
			Caption:   post.Caption,
			ImageURL:  post.ImageURL,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response.Success(c, postResponse, fiber.StatusOK)
}

// GetPostByID godoc
// @Summary Get a post by ID
// @Description Get a post by its unique ID, including the caption, image URL, and timestamps.
// @Tags posts
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} response.GetPostByIDResponse "Successful fetch post response" 
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/posts/{id} [get]
func (h *PostHandler) GetPostByID(c *fiber.Ctx) error {
	id := c.Params("id")

	post, err := h.postService.FetchPostByID(id)
	if err != nil {
		return response.Error(c, "Post not found", fiber.StatusNotFound)
	}

	postResponse := domain.PostResponse{
		ID:        post.ID,
		UserID:    post.UserID,
		Caption:   post.Caption,
		ImageURL:  post.ImageURL,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return response.Success(c, postResponse, fiber.StatusOK)
}

// GetPostsByUserID godoc
// @Summary Get all posts by a specific user
// @Description Get all posts made by a specific user, including the caption, image URL, and timestamps.
// @Tags posts
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} response.GetAllPostsResponse "Successful fetch posts by user response"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/posts/user/{user_id} [get]
func (h *PostHandler) GetPostsByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	posts, err := h.postService.FetchPostsByUserID(userID)
	if err != nil {
		if err == domain.ErrNotFound {
			return response.Error(c, "No posts found for this user", fiber.StatusNotFound)
		}
		return response.Error(c, "Failed to fetch posts", fiber.StatusInternalServerError)
	}

	var postResponse []domain.PostResponse
	for _, post := range posts {
		postResponse = append(postResponse, domain.PostResponse{
			ID:        post.ID,
			UserID:    post.UserID,
			Caption:   post.Caption,
			ImageURL:  post.ImageURL,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response.Success(c, postResponse, fiber.StatusOK)
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post with a caption. Optionally, you can upload an image. If an image is uploaded, its URL will be returned in the response. Requires JWT authentication.
// @Tags posts
// @Accept multipart/form-data
// @Produce json
// @Param caption formData string true "Post caption"
// @Param image formData file false "Post image (optional)"
// @Security BearerAuth
// @Success 201 {object} response.CreatePostResponse "Successful image upload response"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/posts [post]
func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
    if authHeader == "" || len(authHeader) <= len("Bearer ") {
        return response.Error(c.Status(fiber.StatusUnauthorized), "Missing or invalid token")
    }

    token := authHeader[len("Bearer "):]

    ctx := c.Context()
    user, err := h.authService.GetCurrentUser(ctx, token)
    if err != nil {
        return response.Error(c.Status(fiber.StatusUnauthorized), "Unauthorized: "+err.Error())
    }

	caption := c.FormValue("caption")
	if caption == "" {
		return response.ValidationError(c, "Caption is required")
	}

	var imageURL string

	file, err := c.FormFile("image")
	if err == nil {
		uploadDir := "./public/uploads/"
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
				return response.Error(c, "Failed to create upload directory", fiber.StatusInternalServerError)
			}
		}

		sanitizedFileName := sanitizeFileName(file.Filename)

		savePath := uploadDir + sanitizedFileName
		if err := c.SaveFile(file, savePath); err != nil {
			return response.Error(c, "Failed to save image", fiber.StatusInternalServerError)
		}

		imageURL = "http://localhost:8084" + "/uploads/" + sanitizedFileName
	}

	post := domain.Post{
		UserID:   user.ID,
		Caption:  caption,
		ImageURL: imageURL, 
	}

	createdPost, err := h.postService.CreatePost(post)
	if err != nil {
		return response.Error(c, err.Error(), fiber.StatusInternalServerError)
	}

	postResponse := domain.PostResponse{
		ID:        createdPost.ID,
		UserID:    createdPost.UserID,
		Caption:   createdPost.Caption,
		ImageURL:  createdPost.ImageURL,
		CreatedAt: createdPost.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: createdPost.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response.Success(c, postResponse, fiber.StatusCreated)
}

func sanitizeFileName(fileName string) string {
	sanitized := strings.ReplaceAll(fileName, " ", "_")
	sanitized = regexp.MustCompile(`[^a-zA-Z0-9\._-]`).ReplaceAllString(sanitized, "")
	return sanitized
}

// UpdatePost godoc
// @Summary Update an existing post's caption
// @Description Update only the caption of an existing post. Only the post creator is allowed to make this change. Requires JWT authentication.
// @Tags posts
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param request body request.UpdatePostRequest true "Request body with updated caption"
// @Security BearerAuth
// @Success 200 {object} response.UpdatePostResponse "Successful update response"
// @Failure 400 {object} response.ErrorResponse "Bad request
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/posts/{id} [put]
func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
    if authHeader == "" || len(authHeader) <= len("Bearer ") {
        return response.Error(c.Status(fiber.StatusUnauthorized), "Missing or invalid token")
    }

    token := authHeader[len("Bearer "):]

    ctx := c.Context()
    user, err := h.authService.GetCurrentUser(ctx, token)
    if err != nil {
        return response.Error(c.Status(fiber.StatusUnauthorized), "Unauthorized: "+err.Error())
    }

    id := c.Params("id")

    existingPost, err := h.postService.FetchPostByID(id)
    if err != nil {
        return response.Error(c, "Post not found", fiber.StatusNotFound)
    }

    if existingPost.UserID != user.ID {
        return response.Error(c, "Unauthorized to update this post", fiber.StatusUnauthorized)
    }

    var input request.UpdatePostRequest
    if err := c.BodyParser(&input); err != nil {
        return response.ValidationError(c, "Invalid input")
    }

    if input.Caption == "" {
        return response.ValidationError(c, "Caption cannot be empty")
    }

    existingPost.Caption = input.Caption
    updatedPost, err := h.postService.UpdatePost(id, existingPost)
    if err != nil {
        return response.Error(c, "Failed to update post", fiber.StatusInternalServerError)
    }

    postResponse := domain.PostResponse{
        ID:        updatedPost.ID,
        UserID:    updatedPost.UserID,
        Caption:   updatedPost.Caption,
        ImageURL:  updatedPost.ImageURL,
        CreatedAt: updatedPost.CreatedAt.Format("2006-01-02 15:04:05"),
        UpdatedAt: updatedPost.UpdatedAt.Format("2006-01-02 15:04:05"),
    }

    return response.Success(c, postResponse, fiber.StatusOK)
}


// DeletePost godoc
// @Summary Delete a post
// @Description Delete a post by its ID. Only the post creator is allowed to delete it. Requires JWT authentication.
// @Tags posts
// @Param id path string true "Post ID"
// @Security BearerAuth
// @Success 204 {object} response.DeletePostResponse "Successful delete post response"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/posts/{id} [delete]
func (h *PostHandler) DeletePost(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
    if authHeader == "" || len(authHeader) <= len("Bearer ") {
        return response.Error(c.Status(fiber.StatusUnauthorized), "Missing or invalid token")
    }

    token := authHeader[len("Bearer "):]

    ctx := c.Context()
    user, err := h.authService.GetCurrentUser(ctx, token)
    if err != nil {
        return response.Error(c.Status(fiber.StatusUnauthorized), "Unauthorized: "+err.Error())
    }

	id := c.Params("id")

	post, err := h.postService.FetchPostByID(id)
	if err != nil {
		return response.Error(c, "Post not found", fiber.StatusNotFound)
	}

	if post.UserID != user.ID {
		return response.Error(c, "Unauthorized to delete this post", fiber.StatusUnauthorized)
	}

	err = h.postService.DeletePost(id)
	if err != nil {
		return response.Error(c, "Post not found", fiber.StatusNotFound)
	}
	return response.Success(c, fiber.Map{"message": "Post deleted successfully"}, fiber.StatusNoContent)
}

// SearchPosts godoc
// @Summary Search posts
// @Description Search for posts that match a given query, such as a keyword in the caption or content.
// @Tags search
// @Produce json
// @Param query query string true "Search query"
// @Success 200 {array} response.SearchPostsResponse "Successful search response"
// @Failure 400 {object} response.ErrorResponse "Invalid query parameter"
// @Router /api/search/posts [get]
func (h *PostHandler) SearchPosts(c *fiber.Ctx) error {
	query := c.Query("query")

	if query == "" {
		return response.Error(c, "Query parameter is required", fiber.StatusBadRequest)
	}

	posts, err := h.postService.SearchPosts(query)
	if err != nil {
		return response.Error(c, "No posts found", fiber.StatusNotFound)
	}

	var postResponses []domain.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, domain.PostResponse{
			ID:        post.ID,
			UserID:    post.UserID,
			Caption:   post.Caption,
			ImageURL:  post.ImageURL,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return response.Success(c, postResponses)
}