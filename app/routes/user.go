package routes

import (
	"github.com/gorilla/mux"
	controller "../controllers"
)


func UserRouter(r *mux.Router) {
	r.HandleFunc("/signup", controller.Signup).Methods("POST")
	r.HandleFunc("/login", controller.Login).Methods("POST")
}
