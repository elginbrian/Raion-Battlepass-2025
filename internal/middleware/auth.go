package middleware

import (
	"errors"
	"fiber-starter/pkg/response"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func TokenValidationMiddleware(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.Error(c, "Authorization header is required", fiber.StatusUnauthorized)
		}

		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) != 2 {
			return response.Error(c, "Bearer token is required", fiber.StatusUnauthorized)
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method, expected HMAC")
			}
			return []byte(secret), nil
		})

		if err != nil {
			log.Println("Error parsing token:", err)
			return response.Error(c, "Invalid token: "+err.Error(), fiber.StatusUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return response.Error(c, "Invalid or expired token", fiber.StatusUnauthorized)
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now().Add(-time.Minute)) {
				return response.Error(c, "Token has expired", fiber.StatusUnauthorized)
			}
		} else {
			return response.Error(c, "Invalid token claims", fiber.StatusUnauthorized)
		}

		if userID, ok := claims["user_id"].(float64); ok {
			c.Locals("user_id", int(userID))
		} else {
			return response.Error(c, "User ID not found in token", fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}