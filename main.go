package main

import (
	"github.com/Cristian-Infante/REGISTRO_ESTUDIANTES/routes"
	"log"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	puerto := "8082"
	log.Println("Servidor corriendo en http://localhost:" + puerto)
	log.Fatal(http.ListenAndServe(":"+puerto, nil))
}
