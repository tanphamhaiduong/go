package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
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
		ctx := context.WithValue(r.Context(), "TraceID", traceID.String())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func withLogging(next http.Handler) http.Handler {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.TraceLevel)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Method":     r.Method,
			"Host":       r.Host,
			"Header":     r.Header,
			"RemoteAddr": r.RemoteAddr,
			"Proto":      r.Proto,
			"URL":        r.URL,
		}).Info(fmt.Sprintf("Logged connection from %s", r.RemoteAddr))
		next.ServeHTTP(w, r)
	})
}

func withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement Get User and inject to request to here
		log.Info("With Auth")
		next.ServeHTTP(w, r)
	})
}
