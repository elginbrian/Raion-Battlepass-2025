package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}


func Success(c *fiber.Ctx, data interface{}, statusCode ...int) error {
	code := fiber.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return c.Status(code).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func Error(c *fiber.Ctx, message string, statusCode ...int) error {
	code := fiber.StatusInternalServerError
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}

func ValidationError(c *fiber.Ctx, errors interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "fail",
		"errors":  errors,
	})
}

var validate = validator.New()

func ValidateStruct(s interface{}) map[string]string {
    err := validate.Struct(s)

    if err == nil {
        return nil
    }

    errs := make(map[string]string)

    for _, err := range err.(validator.ValidationErrors) {
        errs[err.Field()] = err.Tag()
    }

    return errs
}
