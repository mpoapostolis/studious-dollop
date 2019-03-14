package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// w.Write("hello World")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", sayHello)
	http.ListenAndServe("8080", nil)
	// ?
}
