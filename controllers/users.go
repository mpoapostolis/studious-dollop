package controllers

import "net/http"

// Users controler
func Users(w http.ResponseWriter, r *http.Request) {
	println(r.Body)
}
