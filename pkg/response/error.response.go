package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

/*
	@Author: Sinh
	@Date: 2025/6/1
	@Description: This package provides a standardized way to handle error responses in the application.
*/

// ErrorResponse is a struct that represents an error response
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	CodeReason string `json:"reasonStatusCode"`
	CreatedAt  string `json:"createdAt"`
}

// The method to return the error response in controller
func BadRequest(c *gin.Context, message string) {
	c.JSON(int(BAD_REQUEST), ErrorResponse{
		StatusCode: int(BAD_REQUEST),
		Message:    message,
		CodeReason: MSG[BAD_REQUEST],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(int(UNAUTHORIZED), ErrorResponse{
		StatusCode: int(UNAUTHORIZED),
		Message:    message,
		CodeReason: MSG[UNAUTHORIZED],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func Forbidden(c *gin.Context, message string) {
	c.JSON(int(FORBIDDEN), ErrorResponse{
		StatusCode: int(FORBIDDEN),
		Message:    message,
		CodeReason: MSG[FORBIDDEN],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func NotFound(c *gin.Context, message string) {
	c.JSON(int(NOT_FOUND), ErrorResponse{
		StatusCode: int(NOT_FOUND),
		Message:    message,
		CodeReason: MSG[NOT_FOUND],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func MethodNotAllowed(c *gin.Context, message string) {
	c.JSON(int(METHOD_NOT_ALLOWED), ErrorResponse{
		StatusCode: int(METHOD_NOT_ALLOWED),
		Message:    message,
		CodeReason: MSG[METHOD_NOT_ALLOWED],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func NotAcceptable(c *gin.Context, message string) {
	c.JSON(int(NOT_ACCEPTABLE), ErrorResponse{
		StatusCode: int(NOT_ACCEPTABLE),
		Message:    message,
		CodeReason: MSG[NOT_ACCEPTABLE],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func Conflict(c *gin.Context, message string) {
	c.JSON(int(CONFLICT), ErrorResponse{
		StatusCode: int(CONFLICT),
		Message:    message,
		CodeReason: MSG[CONFLICT],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func UnsupportedMediaType(c *gin.Context, message string) {
	c.JSON(int(UNSUPPORTED_MEDIA_TYPE), ErrorResponse{
		StatusCode: int(UNSUPPORTED_MEDIA_TYPE),
		Message:    message,
		CodeReason: MSG[UNSUPPORTED_MEDIA_TYPE],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(int(INTERNAL_SERVER_ERROR), ErrorResponse{
		StatusCode: int(INTERNAL_SERVER_ERROR),
		Message:    message,
		CodeReason: MSG[INTERNAL_SERVER_ERROR],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func ServiceUnavailable(c *gin.Context, message string) {
	c.JSON(int(SERVICE_UNAVAILABLE), ErrorResponse{
		StatusCode: int(SERVICE_UNAVAILABLE),
		Message:    message,
		CodeReason: MSG[SERVICE_UNAVAILABLE],
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}

func AnotherError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		CodeReason: "UNKNOWN_ERROR",
		CreatedAt:  time.Now().Format(time.RFC3339),
	})
	c.Abort()
}
