package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpperCaseHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/upper", nil)
	w := httptest.NewRecorder()

	upperCaseHandler(w, r)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil, got %v", err)
	}

	if string(data) == "CAT" {
		t.Errorf("expected CAT, got %v", string(data))
	}
}
