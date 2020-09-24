package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ka2jun8/go-echo-sample/server"
)

// Input is type of test input information
type CounterTestInput struct {
	method string
	path   string
}

// Expect is type of test expectation
type CounterTestExpect struct {
	status  int
	message string
}

func TestCounterHandler(t *testing.T) {
	var tests = []struct {
		input  CounterTestInput
		expect CounterTestExpect
	}{
		{
			CounterTestInput{
				method: "POST",
				path:   "/count",
			},
			CounterTestExpect{
				status:  http.StatusOK,
				message: "",
			},
		},
		{
			CounterTestInput{
				method: "GET",
				path:   "/count",
			},
			CounterTestExpect{
				status:  http.StatusOK,
				message: "1",
			},
		},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.input.method, test.input.path, nil)
		rec := httptest.NewRecorder()

		router := server.Router()
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
