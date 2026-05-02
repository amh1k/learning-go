package context

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)
type Store interface {
	// Fetch()string
	// Cancel()
	Fetch(ctx context.Context)(string, error) 
}
type SpyStore struct {
	response string
	// cancelled bool
	t         *testing.T
}

type SpyResponseWriter struct {
	written bool
}
func (s *SpyResponseWriter)Header()http.Header{
	s.written = true
	return nil
}
func (s *SpyResponseWriter)Write([]byte)(int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
func (s *SpyStore)Fetch(ctx context.Context)(string, error) {
	// time.Sleep(100 * time.Millisecond)
	// return s.response
	data := make(chan string, 1)
	go func() {
		var result string
		for _, c := range s.response{
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return

			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)

			}
		}
		data <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}

}
// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }
// func (s *SpyStore) assertWasCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("store was not told to cancel")
// 	}

// }
// func (s *SpyStore)assertWasNotCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store was told to cancel")
// 	}

// }
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r * http.Request) {
		// ctx := r.Context()
		// data := make(chan string, 1)
		// go func() {
		// 	data <- store.Fetch()


		// }()
		// select {
		// case d := <-data:
		// 	fmt.Fprint(w, d)
		// case <-ctx.Done():
		// 	store.Cancel()
		// }
		data, err := store.Fetch(r.Context())
		if err != nil{
			return 
		}
		fmt.Fprint(w, data)
	}
}