package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"encoding/json"
	"os"
	"regexp"
)

var (
	dbName 			 = "test"
	dbUrl  			 = "localhost:27017"
	serverUrl string
	validLocationId string
)



func TestMain(m *testing.M) {
	dbSession := CreateDbSession(dbUrl)
	defer dbSession.Close()

	router := apiRouter(dbSession)

	server := httptest.NewServer(router)
	serverUrl = server.URL
	defer server.Close()

	os.Exit(m.Run())
}

func TestGetLocations(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", serverUrl + "/locations", nil)
	assert.NoError(err)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(err)

	assert.Contains(string(body), "latitude", "%s", "unexpected body returned")
	assert.Contains(string(body), "longitude", "%s", "unexpected body returned")
	assert.Equal(http.StatusOK, res.StatusCode, "%s", "unexpected status code")
}

func TestGetLocationsPaging(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest("GET", serverUrl + "/locations?page=1&per_page=1", nil)
	assert.NoError(err)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(err)

	assert.Contains(string(body), "latitude", "%s", "unexpected body returned")
	assert.Contains(string(body), "longitude", "%s", "unexpected body returned")
	assert.Equal(http.StatusOK, res.StatusCode, "%s", "unexpected status code")

	var jsonObjs interface{}
	errJson := json.Unmarshal([]byte(body), &jsonObjs)
	objSlice, ok := jsonObjs.([]interface{})
	assert.Equal(true, ok, "cannot convert response to JSON object")
	assert.NoError(errJson)
	assert.Equal(1, len(objSlice), "%s", "unexpected number of objects returned")
}

func TestPostInvalidLocation(t *testing.T) {
	assert := assert.New(t)

	userJson := `{"state": "invalid", "response": 500}`
    reader := strings.NewReader(userJson)

	req, err := http.NewRequest("POST", serverUrl + "/locations", reader)
	assert.NoError(err)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(err)

	assert.Contains(string(body), "code", "%s", "unexpected body returned")
	assert.Contains(string(body), "status", "%s", "unexpected body returned")
	assert.Contains(string(body), "Bad Request", "%s", "unexpected body returned")
	assert.Equal(http.StatusBadRequest, res.StatusCode, "%s", "unexpected status code")
}

func TestPostValidLocation(t *testing.T) {
	assert := assert.New(t)

	userJson := `[{"latitude":1.1111,"longitude":2.2222,"altitude":3.3333,"horizontalAccuracy":4.4444,"verticalAccuracy":5.5555,"devicetime":"2016-06-01T07:00:00Z","description":"test location 3"}]`
    reader := strings.NewReader(userJson)

	req, err := http.NewRequest("POST", serverUrl + "/locations", reader)
	assert.NoError(err)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(err)

	assert.Contains(string(body), "id", "%s", "unexpected body returned")
	assert.Equal(http.StatusCreated, res.StatusCode, "%s", "unexpected status code")

	re := regexp.MustCompile("\"id\":\"([a-zA-Z0-9]+)\"")
	validLocationId = re.FindStringSubmatch(string(body))[1]
	assert.NotNil(validLocationId)
}

func TestGetOneLocation(t *testing.T) {
	assert := assert.New(t)

	assert.NotNil(validLocationId)
	req, err := http.NewRequest("GET", serverUrl + "/locations/" + validLocationId, nil)
	assert.NoError(err)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(err)

	assert.Contains(string(body), "id", "%s", "unexpected body returned")
	assert.Contains(string(body), validLocationId, "%s", "unexpected body returned")
	assert.Equal(http.StatusOK, res.StatusCode, "%s", "unexpected status code")
}
