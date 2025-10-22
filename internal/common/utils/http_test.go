package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetHTTPHeaders(t *testing.T) {
	t.Run("sets single header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://example.com", nil)
		headers := map[string]string{
			"Content-Type": "application/json",
		}

		SetHTTPHeaders(req, headers)

		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type header to be application/json, got %s", req.Header.Get("Content-Type"))
		}
	})

	t.Run("sets multiple headers", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://example.com", nil)
		headers := map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer token123",
			"X-Custom":      "custom-value",
		}

		SetHTTPHeaders(req, headers)

		if req.Header.Get("Content-Type") != "application/json" {
			t.Error("Content-Type header not set correctly")
		}
		if req.Header.Get("Authorization") != "Bearer token123" {
			t.Error("Authorization header not set correctly")
		}
		if req.Header.Get("X-Custom") != "custom-value" {
			t.Error("X-Custom header not set correctly")
		}
	})

	t.Run("handles empty headers map", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://example.com", nil)
		headers := map[string]string{}

		SetHTTPHeaders(req, headers)

		// Should not panic and request should still be valid
		if req == nil {
			t.Error("request should not be nil")
		}
	})
}

func TestExecuteHTTPRequest(t *testing.T) {
	t.Run("executes successful request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success response"))
		}))
		defer server.Close()

		req, _ := http.NewRequest("GET", server.URL, nil)
		body := ExecuteHTTPRequest(req)

		if string(body) != "success response" {
			t.Errorf("expected 'success response', got %s", string(body))
		}
	})

	t.Run("panics on non-200 status code", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error message"))
		}))
		defer server.Close()

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for non-200 status code")
			}
		}()

		req, _ := http.NewRequest("GET", server.URL, nil)
		ExecuteHTTPRequest(req)
	})
}
