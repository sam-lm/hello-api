package api_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codeformio/hello-api/api"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(&api.Handler{Message: "Test"})
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	require.Contains(t, string(body), "Test")
}
