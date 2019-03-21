package router

import (
	controllers "../controllers"
	"github.com/gorilla/mux"
)

// NewRouter with handlers
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/users", controllers.Users)
	return r
}
