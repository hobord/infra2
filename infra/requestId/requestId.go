package requestId

import (
	"context"
	"net/http"

	uuid "github.com/google/uuid"
)

type requestKey int

const requestIDKey requestKey = 0

func newContextWithRequestID(ctx context.Context, req *http.Request) context.Context {
	uuid := uuid.New()
	return context.WithValue(ctx, requestIDKey, uuid.String())
}

// RequestIDFromContext is get the request id from context
func RequestIDFromContext(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

// RequestIDHandler provide a midlleware
func RequestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := newContextWithRequestID(req.Context(), req)
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
