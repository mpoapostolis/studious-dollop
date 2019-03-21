package main

import (
	"log"
	"net/http"
)

func main() {
	println("serve and run at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
