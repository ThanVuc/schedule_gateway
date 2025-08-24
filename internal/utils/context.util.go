package utils

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func WithRequestID(ctx context.Context, requestID string) context.Context {
	if requestID == "" {
		return ctx
	}
	md := metadata.Pairs("x-request-id", requestID)
	return metadata.NewOutgoingContext(ctx, md)
}

func GetHttpOnlyCookie(name, value string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   "",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	return cookie
}
