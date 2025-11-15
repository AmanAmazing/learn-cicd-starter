package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name          string
		headerValue   string
		expectedKey   string
		expectedError bool
	}{
		{
			name:          "Successful Authorization Header",
			headerValue:   "ApiKey as382as",
			expectedKey:   "Authorization",
			expectedError: false,
		},
		{
			name:          "API Key Blank",
			headerValue:   "ApiKey",
			expectedKey:   "Authorization",
			expectedError: true,
		},
		{
			name:          "Authorization Header not set",
			headerValue:   "",
			expectedKey:   "",
			expectedError: true,
		},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(http.MethodPost, "/auth", nil)
		req.Header.Set(tc.expectedKey, tc.headerValue)

		gotString, gotError := GetAPIKey(req.Header)
		if tc.expectedError == true {
			if gotError == nil {
				t.Fatalf("Expected an error here but did not receive one: %s", gotString)
			}
		}

		if tc.expectedError == false && gotError != nil {
			t.Fatal("Didn't expect an error here but got one")
		}

	}

}
