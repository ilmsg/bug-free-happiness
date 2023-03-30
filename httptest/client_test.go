package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClientUpperCase(t *testing.T) {
	expected := "dummy data"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, expected)
	}))
	defer srv.Close()

	c := NewClient(srv.URL)
	res, err := c.UpperCaseHandler("anything")
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}

	res = strings.TrimSpace(res)
	if res != expected {
		t.Errorf("expected res to be %s, got %s", expected, res)
	}

}
