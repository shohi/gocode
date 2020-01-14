package main

// refer, support Python etc non-executable scripts on Windows
// https://github.com/golang/go/issues/18420
import (
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
	"os"
)

func handleCGI(w http.ResponseWriter, r *http.Request) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("get current directory error=%s\n", err)
	}
	handler := new(cgi.Handler)
	handler.Path = currentDir + r.URL.Path
	handler.Dir = ""
	handler.ServeHTTP(w, r)
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("get current directory error=%s\n", err)
	}
	http.Handle("/", http.FileServer(http.Dir(currentDir)))
	http.HandleFunc("/cgi-bin/", handleCGI)
	fmt.Printf("server listening 127.0.0.1:8000 ...\n")
	http.ListenAndServe(":8000", nil)
}
