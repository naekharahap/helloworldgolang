package main

import (
	"helloworld/helloworld"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(helloworld.HelloWorld()))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
