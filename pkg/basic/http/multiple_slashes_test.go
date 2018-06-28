package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMultipleSlash(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v", r)
		if strings.EqualFold(r.Method, "put") {
			var b []byte
			var err error
			if r.Body != nil {
				defer r.Body.Close()
				b, err = ioutil.ReadAll(r.Body)
				if err != nil {
					log.Printf("error info ==> %v", err)
				}
			}
			log.Printf("contet ==> %v", string(b))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HaHa"))
	})

	s := httptest.NewServer(handler)
	urlFormat := "%s/%s"
	url := fmt.Sprintf(urlFormat, s.URL, "test//data/1")
	log.Printf("request url ==> %v\n", url)
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader("hello"))

	if err != nil {
		log.Printf("create request error: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	log.Println(res, err)

	// test external
	client := &http.Client{}
	url = "http://54.169.170.244:9002/pc3/user//data/1"
	req, _ = http.NewRequest(http.MethodPut, url, strings.NewReader("hello"))
	res, err = client.Do(req)
	log.Println(res, err)

	// test search engine
	url = "http://www.baidu.com/user//data/1"
	req, _ = http.NewRequest(http.MethodPut, url, strings.NewReader("hello"))
	res, err = client.Do(req)
	log.Println(res, err)

	// test local engine
	url = "http://localhost:9080/user//data/1"
	req, _ = http.NewRequest(http.MethodPut, url, strings.NewReader("hello"))
	res, err = client.Do(req)
	log.Println(res, err)
}

func TestMultipleSlash301(t *testing.T) {
	client := &http.Client{}

	// test local engine
	url := "http://54.169.170.244:9002/pc3/user//data/4"
	req, _ := http.NewRequest(http.MethodPut, url, strings.NewReader("hello"))
	res, err := client.Do(req)
	log.Println(res, err)
}
