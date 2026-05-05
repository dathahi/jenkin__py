package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler := setupRoutes()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	ct := rec.Header().Get("Content-Type")
	if !strings.Contains(ct, "application/json") {
		t.Errorf("expected Content-Type application/json, got %s", ct)
	}

	var body map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}

	if body["status"] != "ok" {
		t.Errorf("expected status ok, got %s", body["status"])
	}
}

func TestRootEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler := setupRoutes()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	ct := rec.Header().Get("Content-Type")
	if !strings.Contains(ct, "text/html") {
		t.Errorf("expected Content-Type text/html, got %s", ct)
	}

	body := rec.Body.String()
	required := []string{"Go Demo", "Jenkins Learning", "Version 1.0.0", "Hostname"}
	for _, s := range required {
		if !strings.Contains(body, s) {
			t.Errorf("expected body to contain %q", s)
		}
	}
}

func TestNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	rec := httptest.NewRecorder()

	handler := setupRoutes()
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rec.Code)
	}
}

func TestHealthResponseFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler := setupRoutes()
	handler.ServeHTTP(rec, req)

	var result map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 key in response, got %d", len(result))
	}

	_, exists := result["status"]
	if !exists {
		t.Error("expected 'status' key in response")
	}
}

func TestEnvVarsNotSet(t *testing.T) {
	os.Unsetenv("API_KEY")
	os.Unsetenv("DB_PASSWORD")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler := setupRoutes()
	handler.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "API_KEY:</span>") || !strings.Contains(body, "(not set)") {
		t.Error("expected env vars to show (not set) when not set")
	}
}

func TestEnvVarsSet(t *testing.T) {
	os.Setenv("API_KEY", "test-api-key")
	os.Setenv("DB_PASSWORD", "test-db-password")
	defer os.Unsetenv("API_KEY")
	defer os.Unsetenv("DB_PASSWORD")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler := setupRoutes()
	handler.ServeHTTP(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "API_KEY:</span>") || !strings.Contains(body, "****") {
		t.Error("expected masked credentials when env vars are set")
	}
	if strings.Contains(body, "test-api-key") || strings.Contains(body, "test-db-password") {
		t.Error("actual credential values should not appear in response")
	}
}