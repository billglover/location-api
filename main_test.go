package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	dbName = "test"
	dbUrl  = "localhost:27017"
)

var tests = []struct {
	Method       string
	Path         string
	Body         io.Reader
	BodyContains string
	Status       int
}{
	{
		Method:       "GET",
		Path:         "/locations",
		BodyContains: "",
		Status:       http.StatusNotFound,
	},
	{
		Method:       "POST",
		Path:         "/locations",
		Body:         strings.NewReader(`{""}`),
		BodyContains: "",
		Status:       http.StatusNotImplemented,
	},
	{
		Method:       "GET",
		Path:         "/locations/574cb30f4bf4c8f0c6a056e8",
		BodyContains: "",
		Status:       http.StatusOK,
	},
}

func TestAll(t *testing.T) {
	assert := assert.New(t)

	// create a DB Session
	// TODO: this needs to be able to take a database name
	dbSession := CreateDbSession(dbUrl)
	defer dbSession.Close()

	// create a Router
	router := apiRouter(dbSession)

	// create a test server
	server := httptest.NewServer(router)
	defer server.Close()

	// execute our tests
	for _, test := range tests {

		// create and execute an HTTP request
		r, err := http.NewRequest(test.Method, server.URL+test.Path, test.Body)
		assert.NoError(err)
		response, err := http.DefaultClient.Do(r)
		assert.NoError(err)

		// extract the body of the response
		actualBody, err := ioutil.ReadAll(response.Body)
		assert.NoError(err)

		// make assertions

		assert.Contains(string(actualBody), test.BodyContains, "%s %s %s", test.Method, test.Path, "\n\tunexpected body returned")
		assert.Equal(test.Status, response.StatusCode, "%s %s %s", test.Method, test.Path, "\n\tunexpected status code in response")
	}
}
