package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	logger "./app/lib"
	routes "./app/routes"
	models "./app/models"
)


func main()  {
	router := mux.NewRouter()
	models.Connect()
	routes.UserRouter(router)
	router.Use(logger.Log)
	log.Printf("Server listning on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
