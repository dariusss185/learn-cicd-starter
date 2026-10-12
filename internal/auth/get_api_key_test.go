package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		wantAPIKey string
		wantErr    bool
	}{
		{
			name:       "valid API key",
			authHeader: "ApiKey abc123",
			wantAPIKey: "abc123",
		},
		{
			name:    "missing authorization header",
			wantErr: true,
		},
		{
			name:       "missing API key",
			authHeader: "ApiKey",
			wantErr:    true,
		},
		{
			name:       "incorrect authentication scheme",
			authHeader: "Bearer abc123",
			wantErr:    true,
		},
		{
			name:       "empty API key",
			authHeader: "ApiKey ",
			wantAPIKey: "",
		},
		{
			name:       "extra spaces",
			authHeader: "ApiKey abc123 extra",
			wantAPIKey: "abc123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			got, err := GetAPIKey(headers)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.wantAPIKey {
				t.Errorf("expected API key %q, got %q", tt.wantAPIKey, got)
			}
		})
	}
}
