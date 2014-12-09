package basic

import (
	"log"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestNoAuthorizationHeader(t *testing.T) {

	ts := createServer(trulyValidator)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 401, res.StatusCode)
}

func TestEmptyAuthorizationHeader(t *testing.T) {

	ts := createServer(trulyValidator)
	defer ts.Close()

	client := &http.Client{}

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 401, res.StatusCode)
}

func TestTruelyAuthorizationHeader(t *testing.T) {

	ts := createServer(trulyValidator)
	defer ts.Close()

	client := &http.Client{}

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("foo", "bar")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 204, res.StatusCode)
}

func TestFalsyAuthorizationHeader(t *testing.T) {

	ts := createServer(falsyValidator)
	defer ts.Close()

	client := &http.Client{}

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("foo", "bar")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 401, res.StatusCode)
}

func createServer(validator func(username, password string) bool) *httptest.Server {
	auth := NewAuthenticator(validator, "testServer")

	appHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})

	return httptest.NewServer(auth.Wrap(appHandler))
}

func trulyValidator(username, password string) bool {
  return true
}

func falsyValidator(username, password string) bool {
  return false
}
