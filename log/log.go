package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type valueKey int

const requestIDKey valueKey = 42

// Println prints request Id from context
func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		log.Println("Could not find request ID in req")
		return
	}

	log.Printf("[%d] %s", id, msg)
}

// Decorate wraps the handlerfunc
func Decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := rand.Int63()
		ctx = context.WithValue(ctx, requestIDKey, id)

		f(w, r.WithContext(ctx))
	}
}
