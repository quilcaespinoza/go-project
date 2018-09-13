package routes

import (
	"github.com/golang-es/prjecomments/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetUserRouter es la ruta para usuarios
func SetUserRouter(router *mux.Router) {
	prefix := "/api/user"

	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)

}
