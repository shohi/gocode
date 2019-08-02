package httptest_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readAndCloseBody(resp *http.Response) (string, error) {
	if resp == nil {
		return "", nil
	}

	if resp.Body == nil {
		return "", nil
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	return string(content), err
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("default handler"))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("data handler"))
}

func redictHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/data", http.StatusFound)
}

func TestServer_Mux(t *testing.T) {
	assert := assert.New(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/data", dataHandler)
	mux.HandleFunc("/", defaultHandler)

	srv := httptest.NewServer(mux)

	// Case 1
	resp, err := http.Get(srv.URL + "/")
	assert.Nil(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	content, err := readAndCloseBody(resp)
	assert.Nil(err)
	assert.Contains(content, "default")

	// Case 2
	resp, err = http.Get(srv.URL + "/data")
	assert.Nil(err)

	content, err = readAndCloseBody(resp)
	assert.Nil(err)
	assert.Contains(content, "data")
}

func TestServer_302(t *testing.T) {
	assert := assert.New(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/data", dataHandler)
	mux.HandleFunc("/302", redictHandler)

	srv := httptest.NewServer(mux)

	// NOTE: standard golang http client follows 302 redirects.
	// Case - test 302
	resp, err := http.Get(srv.URL + "/302")
	assert.Nil(err)

	content, err := readAndCloseBody(resp)
	assert.Nil(err)
	assert.Contains(content, "data")
}
