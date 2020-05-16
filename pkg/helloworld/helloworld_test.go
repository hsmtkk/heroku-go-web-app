package helloworld_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsmtkk/heroku-go-web-app/pkg/helloworld"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(helloworld.Handler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.Nil(t, err, "should be nil")
	defer res.Body.Close()

	want := []byte("Hello, World!")
	got, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, want, got, "should be equal")
}
