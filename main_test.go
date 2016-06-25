package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"encoding/json"
)

var (
	dbName = "test"
	dbUrl  = "localhost:27017"
)

var tests = []struct {
	Method       string
	Path         string
	Body         io.Reader
	BodyContains []string
	BodyCount	 int
	Status       int
}{
	{
		Method:       "GET",
		Path:         "/locations",
		BodyContains: []string{"latitude", "longitude"},
		Status:       http.StatusOK,
	},
	{
		Method:       "GET",
		Path:         "/locations?page=1&per_page=1",
		BodyContains: []string{""},
		BodyCount:	  1,
		Status:       http.StatusOK,
	},
	{
		Method:       "GET",
		Path:         "/locations?page=2&per_page=1",
		BodyContains: []string{""},
		BodyCount:	  1,
		Status:       http.StatusOK,
	},
	{
		Method:       "GET",
		Path:         "/locations?per_page=2",
		BodyContains: []string{""},
		BodyCount:	  2,
		Status:       http.StatusOK,
	},
	{
		Method:       "POST",
		Path:         "/locations",
		Body:         strings.NewReader(`{"name": "dave"}`),
		BodyContains: []string{""},
		Status:       http.StatusBadRequest,
	},
	{
		Method:       "POST",
		Path:         "/locations",
		Body:         strings.NewReader(`[{"name": "dave"}]`),
		BodyContains: []string{""},
		Status:       http.StatusBadRequest,
	},
	{
		Method:       "POST",
		Path:         "/locations",
		Body:         strings.NewReader(`[{"latitude":1.1111,"longitude":2.2222,"altitude":3.3333,"horizontalAccuracy":4.4444,"verticalAccuracy":5.5555,"devicetime":"2016-06-01T07:00:00Z","description":"test location 3"}]`),
		BodyContains: []string{"\"latitude\":1.1111,\"longitude\":2.2222,\"altitude\":3.3333,\"horizontalAccuracy\":4.4444,\"verticalAccuracy\":5.5555,\"devicetime\":\"2016-06-01T07:00:00Z\",\"description\":\"test location 3\""},
		Status:       http.StatusCreated,
	},
	{
		Method:       "GET",
		Path:         "/locations/574de23b5f810df11cad3498",
		BodyContains: []string{"\"id\":\"574de23b5f810df11cad3498\""},
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

		// convert body into an array of structs
		if test.BodyCount > 0 {
			var jsonObjs interface{}
	 		errJson := json.Unmarshal([]byte(actualBody), &jsonObjs)
	 		objSlice, ok := jsonObjs.([]interface{})
	 		assert.Equal(true, ok, "cannot convert response to JSON object")
			assert.NoError(errJson)
			assert.Equal(test.BodyCount, len(objSlice), "%s %s %s", test.Method, test.Path, "\n\tunexpected number of objects returned")
		}

		// make assertions
		assert.Equal(test.Status, response.StatusCode, "%s %s %s", test.Method, test.Path, "\n\tunexpected status code in response")
		for _, bodyContents := range test.BodyContains {
			assert.Contains(string(actualBody), bodyContents, "%s %s %s", test.Method, test.Path, "\n\tunexpected body returned")
		}
	}
}
