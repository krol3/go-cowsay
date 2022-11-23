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

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homePage)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Welcome to the HomePage CowSay!`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
