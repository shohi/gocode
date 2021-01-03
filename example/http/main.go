package main

import "net/http"

func defaultHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func main() {

	err := http.ListenAndServe(":8090",
		http.HandlerFunc(defaultHandleFunc))
	if err != nil {
		panic(err)
	}
}
