package main

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/tanphamhaiduong/go/common/logger"
	"github.com/tanphamhaiduong/go/delta/internal/models"
)

// middleware provides a convenient mechanism for filtering HTTP requests
// entering the application. It returns a new handler which may perform various
// operations and should finish by calling the next HTTP handler.
type middleware func(next http.HandlerFunc) http.HandlerFunc

// chainMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
func chainMiddleware(mw ...middleware) middleware {
	return func(final http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(w, r)
		}
	}
}

func withTimeout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Duration(60*time.Second))
		defer cancel()
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func withTraceID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID, _ := uuid.NewRandom()
		var traceIDKey models.ContextKey = "traceId"
		ctx := context.WithValue(r.Context(), traceIDKey, traceID.String())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func withLogging(next http.Handler) http.Handler {
	// Only log the warning severity or above.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var traceIDKey models.ContextKey = "traceId"
		logger.WithFields(logger.Fields{
			"traceID":       r.Context().Value(traceIDKey),
			"host":          r.Host,
			"remoteAddr":    r.RemoteAddr,
			"method":        r.Method,
			"requestUri":    r.RequestURI,
			"proto":         r.Proto,
			"connection":    r.Header.Get("Connection"),
			"contentLength": r.ContentLength,
			"contentType":   r.Header.Get("Content-Type"),
			"userAgent":     r.Header.Get("User-Agent"),
			"url":           r.URL,
		}).Infof("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
		start := time.Now()
		duration := time.Now().Sub(start)
		logger.WithFields(logger.Fields{
			"traceId":  r.Context().Value(traceIDKey),
			"Duration": duration,
		}).Infof("Logged connection from %s", r.RemoteAddr)
	})
}

func withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			secretKey        = os.Getenv("SECRET_KEY")
			jwtKey           = []byte(secretKey)
			authorization    = r.Header.Get("Authorization")
			tokenString      = strings.ReplaceAll(authorization, "Bearer ", "")
			isDevelopmentEnv = secretKey == ""
		)
		if isDevelopmentEnv {
			jwtKey = []byte("test")
		}
		claims := &models.Claims{}
		_, _ = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		var claimsKey models.ContextKey = "claims"
		ctx := context.WithValue(r.Context(), claimsKey, claims)
		r = r.WithContext(ctx)
		// Implement Get User and inject to request to here
		next.ServeHTTP(w, r)
	})
}
