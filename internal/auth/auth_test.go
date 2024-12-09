package auth

import (
	"net/http"
	"testing"
)

func TestApiKeyExists(t *testing.T) {
	apiKey := "api-key-123"
	h := http.Header{}
	h.Add("Authorization", "ApiKey "+apiKey)
	key, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("expected no error")
	}

	if key != apiKey {
		t.Fatalf("expected: %v, got: %v", apiKey, key)
	}
}

func TestNoApiKey(t *testing.T) {
	h := http.Header{}
	h.Add("Authorization", "ApiKey")
	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatalf("expected error to be thrown")
	}
	if err.Error() != "malformed authorization header" {
		t.Fatalf("Expected: malformed authorization header but got: %v", err.Error())
	}
}
