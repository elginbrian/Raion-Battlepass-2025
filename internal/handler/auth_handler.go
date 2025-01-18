package handler

import (
	"fiber-starter/internal/service"
	"fiber-starter/pkg/request"
	"fiber-starter/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// @Summary Register a new user
// @Description Create a new account by providing a username, email, and password. The system checks if the details are valid and returns a success message if registration is successful.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.UserRegistrationRequest true "User registration details"
// @Success 201 {object} response.RegisterResponse "Successful registration response"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req request.UserRegistrationRequest

	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request format")
	}

	validate := validator.New()
	if validationErrs := validate.Struct(req); validationErrs != nil {
		return response.ValidationError(c, validationErrs.Error())
	}

	if err := h.authService.Register(req.Username, req.Email, req.Password); err != nil {
		return response.Error(c, err.Error())
	}

	return response.Success(c.Status(fiber.StatusCreated), response.RegisterData{
		Message: "User registered successfully",
	})
}

// @Summary Log in a user
// @Description Log in to your account by providing your email and password. If the details are correct, you will receive a JWT token to use for secure access to other endpoints.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body request.UserLoginRequest true "User login details"
// @Success 201 {object} response.LoginResponse "Successful registration response"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req request.UserLoginRequest

	if err := c.BodyParser(&req); err != nil {
		return response.ValidationError(c, "Invalid request format")
	}

	validate := validator.New()
	if validationErrs := validate.Struct(req); validationErrs != nil {
		return response.ValidationError(c, validationErrs.Error())
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		return response.Error(c.Status(fiber.StatusUnauthorized), err.Error())
	}

	return response.Success(c, response.LoginData{
		Token: "Bearer " + token,
	})
}

// @Summary Get current user info
// @Description Retrieve details about the logged-in user. You must include your JWT token in the Authorization header to access this information.
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.GetCurrentUserResponse "User information retrieved successfully"
// @Failure 401 {object} response.ErrorResponse "Unauthorized or invalid token"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/auth/current-user [get]
func (h *AuthHandler) GetUserInfo(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" || len(authHeader) <= len("Bearer ") {
        return response.Error(c.Status(fiber.StatusUnauthorized), "Missing or invalid token")
    }

    token := authHeader[len("Bearer "):]

    ctx := c.Context()
    user, err := h.authService.GetCurrentUser(ctx, token)
    if err != nil {
        return response.Error(c.Status(fiber.StatusUnauthorized), err.Error())
    }

    return response.Success(c, response.User{
        ID:        user.ID,
        Username:  user.Name,
        Email:     user.Email,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
    })
}

// @Summary Change your password
// @Description Update your password securely. You need to be logged in and provide your old password along with the new one. Include your JWT token in the Authorization header.
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param data body request.ChangePasswordRequest true "Change Password Request"
// @Success 200 {object} response.ChangePasswordData "Password changed successfully"
// @Failure 400 {object} response.ErrorResponse "Validation error or invalid request format"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/auth/change-password [put]
func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
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

    var req request.ChangePasswordRequest
    if err := c.BodyParser(&req); err != nil {
        return response.ValidationError(c, "Invalid request format")
    }

    validate := validator.New()
    if validationErrs := validate.Struct(req); validationErrs != nil {
        return response.ValidationError(c, validationErrs.Error())
    }

    if err := h.authService.ChangePassword(user.ID, req.OldPassword, req.NewPassword); err != nil {
        return response.Error(c, err.Error())
    }

    return response.Success(c, response.ChangePasswordData{
        Message: "Password changed successfully",
    })
}