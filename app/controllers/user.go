package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)
import  "../models"


func Login (w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	user := models.User{}
	models.DB.First(&user)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(make(map[string] interface { "user": user}))
}

func Signup(w http.ResponseWriter, r *http.Request)  {
	
}