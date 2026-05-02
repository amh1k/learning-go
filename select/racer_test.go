package selectfolder

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)
func TestRacer(t *testing.T) {
	t.Run("comparing speed of servers, returns url of faster one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()
	slowURL := slowServer.URL
	fastURL := fastServer.URL
	want := fastURL
	got, _ := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	})
	t.Run("returns an error if server does not respond within 10s", func( t* testing.T) {
		serverA := makeDelayedServer(5 * time.Second)
		serverB := makeDelayedServer(5 * time.Second)
		defer serverA.Close()
		defer serverB.Close()
		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 5 *time.Millisecond)
		if err == nil {
			t.Error("expected error but didnt get one")
		}
	})
	

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}