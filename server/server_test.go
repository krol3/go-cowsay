package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Tests struct {
	name          string
	request       *http.Request
	httpCode      int
	body          string
	expectedError error
}

func TestGetServer(t *testing.T) {
	tests := []Tests{
		{
			name:          "TestGetServer",
			request:       httptest.NewRequest(http.MethodGet, "/cow?key=abc", nil),
			httpCode:      200,
			body:          "abc",
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := httptest.NewRecorder()
			cowPage(response, test.request)
			if response.Code != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, response.Code)
			}

			if !strings.Contains(response.Body.String(), test.body) {
				t.Errorf(
					`response body "%s" does not contain "%s"`,
					response.Body.String(), test.body,
				)
			}

		})
	}
}
