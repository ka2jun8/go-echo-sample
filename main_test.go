package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type Input struct {
	method string
	path   string
}

type Expect struct {
	status  int
	message string
}

func TestHealthzHandler(t *testing.T) {
	var tests = []struct {
		input  Input
		expect Expect
	}{
		{
			Input{
				method: "GET",
				path:   "/healthz",
			},
			Expect{
				status:  http.StatusOK,
				message: "ok",
			},
		},
		{
			Input{
				method: "POST",
				path:   "/count",
			},
			Expect{
				status:  http.StatusOK,
				message: "",
			},
		},
		{
			Input{
				method: "GET",
				path:   "/count",
			},
			Expect{
				status:  http.StatusOK,
				message: "1",
			},
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.input.method, test.input.path, nil)
		rec := httptest.NewRecorder()

		router := Router()
		router.ServeHTTP(rec, req)

		status := test.expect.status
		message := test.expect.message
		actualStatus := rec.Code
		actualMessage := rec.Body.String()

		if status != actualStatus {
			t.Errorf("expected: %v, result: %q", status, actualStatus)
		}
		if message != actualMessage {
			t.Errorf("expected: %v, result: %q", message, actualMessage)
		}
	}
}
