package handlers_test

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/matryer/respond"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
		Status:       http.StatusNotImplemented,
	},
	{
		Method:       "POST",
		Path:         "/locations",
		Body:         strings.NewReader(`{""}`),
		BodyContains: "",
		Status:       http.StatusCreated,
	},
	{
		Method:       "GET",
		Path:         "/locations/1234",
		BodyContains: "",
		Status:       http.StatusOK,
	},
}

func TestAll(t *testing.T) {
	assert := assert.New(t)
	server := httptest.NewServer(&myhandler{})
	defer server.Close()
	for _, test := range tests {
		r, err := http.NewRequest(test.Method, server.URL+test.Path, test.Body)
		assert.NoError(err)
		// call handler
		response, err := http.DefaultClient.Do(r)
		assert.NoError(err)
		actualBody, err := ioutil.ReadAll(response.Body)
		assert.NoError(err)
		// make assertions
		assert.Contains(string(actualBody), test.BodyContains, "body")
		assert.Equal(test.Status, response.StatusCode, "status code")
	}
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respond.WithStatus(w, r, http.StatusOK)
}

type myhandler struct{}

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "TestAll", F: TestAll})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
