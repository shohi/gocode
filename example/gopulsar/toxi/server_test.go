package toxi

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("===> receive request: %v\n",
		req.URL.Path)

	w.Write([]byte("OK"))
}

func newTestServer(port int) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(mainHandler),
	}

	go srv.ListenAndServe()

	return srv
}
