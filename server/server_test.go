package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *http.ResponseWriter
	expectedError error
}

func TestGetServer(t *testing.T) {
	tests := []Tests{
name: "TestGetServer",
server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
},
response:    httptest.ResponseWriter(),){
}),
expectedError: nil,
}
