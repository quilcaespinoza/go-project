package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/golang-es/prjecomments/migration"
	"github.com/golang-es/prjecomments/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()

	if migrate == "yes" {
		log.Println("Comenzó la migración")
		migration.Migrate()
		log.Println("Terminó la migración")
	}
	// Inicia las rutas
	router := routes.InitRoutes()

	//Inicia los midleware
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("Iniciado el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Fatal("Final del programa")
}
