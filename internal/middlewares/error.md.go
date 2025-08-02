package middlewares

import (
	"net/http"
	"schedule_gateway/global"
	"schedule_gateway/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

/*
@Author: Sinh
@Date:
@Description: Global error handler middleware for the application.
*/
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logger := global.Logger
				requestIdVal, _ := c.Get("requestId")
				requestId, _ := requestIdVal.(string)
				switch e := r.(type) {
				case response.ErrorResponse:
					// Needing more details ? -> add debug.Stack() to the logger
					if e.StatusCode >= 500 {
						logger.Error("Internal Server Error", requestId)
					} else {
						logger.Error("Internal Server Error", requestId)
					}

					c.JSON(e.StatusCode, e)
				default:
					c.JSON(http.StatusInternalServerError, response.ErrorResponse{
						StatusCode: http.StatusInternalServerError,
						Message:    "Internal Server Error",
						CodeReason: "INTERNAL_SERVER_ERROR",
						CreatedAt:  time.Now().Format(time.RFC3339),
					})
				}
			}

			c.Abort()
		}()
		c.Next()
	}
}
