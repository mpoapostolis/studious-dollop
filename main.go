package main

import (
	"log"
	"net/http"

	router "./router"
)

func main() {
	println("serve and run at Port 8080")
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
