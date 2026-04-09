package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func EnrichContext(ctx context.Context, c *gin.Context) context.Context {
	requestID := c.GetString("request_id")
	userID := c.GetString("user_id")
	groupId := c.Param("group_id")
	origin := c.GetHeader("Origin")

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}

	md = md.Copy()

	md.Set("x-request-id", requestID)
	md.Set("x-user-id", userID)
	if groupId != "" {
		md.Set("x-group-id", groupId)
	}
	baseURL := GetBaseURL(c)
	md.Set("x-base-url", baseURL)
	if origin != "" {
		md.Set("x-origin", origin)
	}

	return metadata.NewOutgoingContext(ctx, md)
}

func GetHttpOnlyCookie(name, value string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   "",
		MaxAge:   7 * 24 * 60 * 60,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	return cookie
}

func ClearCookie(name string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Domain:   "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	return cookie
}

func GetBaseURL(ctx *gin.Context) string {
	scheme := ctx.GetHeader("X-Forwarded-Proto")
	if scheme == "" {
		if ctx.Request.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}

	return fmt.Sprintf("%s://%s", scheme, ctx.Request.Host)
}
