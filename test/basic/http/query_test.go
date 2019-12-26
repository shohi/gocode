package http_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryParameter(t *testing.T) {
	assert := assert.New(t)

	handler := func(w http.ResponseWriter, r *http.Request) {
		addr := r.URL.Query().Get("addr")

		log.Printf("====> addr: [%v]", addr)

		if addr != "http://localhost:6001" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	r, _ := http.NewRequest("GET", "/hello?addr=http://localhost:6001", nil)
	w := httptest.NewRecorder()

	handler(w, r)

	assert.Equal(200, w.Code)
}
