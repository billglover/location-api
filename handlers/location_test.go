package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAll(t *testing.T) {
	var tests = []struct {
		Fn func(rw http.ResponseWriter, req *http.Request)
		S  int    // status
		B  string //body
	}{{
		Fn: func(rw http.ResponseWriter, req *http.Request) {
			rw.Write([]byte("Hello World!"))
		},
		S: 200,
		B: "Hello World!",
	}}

	for _, test := range tests {
		rw := httptest.NewRecorder()
		test.Fn(rw, nil)
		t.Log(rw.Code)
		t.Log(rw.Body.String())
		if test.S != rw.Code {
			t.Fail()
		}
		if test.B != rw.Body.String() {
			t.Fail()
		}
	}
}
