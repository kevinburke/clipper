package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kevinburke/nacl"
)

func TestServer(t *testing.T) {
	mux := NewServeMux()
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("GET /: got code %d, want 200", w.Code)
	}
	if body := w.Body.String(); !strings.Contains(body, "Clipper CSV Download") {
		t.Errorf("GET /: expected 'Clipper CSV Download' in body, got %s", body)
	}
}

func BenchmarkHomepage(b *testing.B) {
	mux := NewServeMux()
	s := httptest.NewServer(mux)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, err := http.Get(s.URL)
		if err != nil {
			b.Fatal(err)
		}
		if res.StatusCode != 200 {
			b.Fatalf("GET /: expected code 200, got %d", res.StatusCode)
		}
		n, err := ioutil.ReadAll(res.Body)
		if err != nil {
			b.Fatal(err)
		}
		b.SetBytes(int64(len(n)))
		res.Body.Close()
	}
}

var sink string

// Just curious about how fast secretbox encryption is
func BenchmarkOpaque(b *testing.B) {
	key := nacl.NewKey()
	secret := "this is an average length message, about sixty characters long."
	b.SetBytes(int64(len(secret)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = opaque(secret, key)
	}
}
