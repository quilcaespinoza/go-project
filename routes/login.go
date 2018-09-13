package routes

import (
	"github.com/golang-es/prjecomments/controllers"
	"github.com/gorilla/mux"
)

// SetLoginRouter es el router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
