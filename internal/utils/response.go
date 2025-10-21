package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response adalah format standar semua API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// RespondSuccess mengembalikan respon sukses dengan data
func RespondSuccess(c echo.Context, status int, message string, data interface{}) error {
	res := Response{
		Success: true,
		Message: message,
		Data:    data,
	}
	return c.JSON(status, res)
}

// RespondCreated shortcut untuk respon 201
func RespondCreated(c echo.Context, message string, data interface{}) error {
	return RespondSuccess(c, http.StatusCreated, message, data)
}

// RespondError mengembalikan respon error
func RespondError(c echo.Context, status int, message string, err interface{}) error {
	res := Response{
		Success: false,
		Message: message,
		Error:   err,
	}
	return c.JSON(status, res)
}
