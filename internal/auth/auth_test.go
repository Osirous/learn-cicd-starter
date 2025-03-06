package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test cases
	testCases := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError bool
	}{
		{
			name:          "Valid API Key",
			headers:       http.Header{"Authorization": []string{"ApiKey test-key-123"}},
			expectedKey:   "test-key-123",
			expectedError: false,
		},
		{
			name:          "Missing Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: true,
		},
		{
			name:          "Malformed Authorization Header",
			headers:       http.Header{"Authorization": []string{"Bearer invalid-format"}},
			expectedKey:   "",
			expectedError: true,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			if tc.expectedError && err == nil {
				t.Fatalf("expected an error but got none")
			}

			if !tc.expectedError && err != nil {
				t.Fatalf("expected no error but got: %v", err)
			}

			if key != tc.expectedKey {
				t.Fatalf("expected key %q, got %q", tc.expectedKey, key)
			}
		})
	}
}
