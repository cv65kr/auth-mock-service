package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTokenGenerateHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/token", strings.NewReader("{ \"username\": \"user-admin\", \"password\": \"password\" }"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	srv := NewMockServer()
	handler := http.HandlerFunc(srv.tokenGenerateHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response tokenResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.Token == "" {
		t.Error("handler returned no token")
	}

	if response.Username != "user" {
		t.Error("handler returned invalid username")
	}

	if response.Role != "ADMIN" {
		t.Error("handler returned invalid role")
	}
}

func TestPublicKeyHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/public-key", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	srv := NewMockServer()

	handler := http.HandlerFunc(srv.publicKeyHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response publicKeyResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if response.PublicKey == "" {
		t.Error("handler returned no public key")
	}
}