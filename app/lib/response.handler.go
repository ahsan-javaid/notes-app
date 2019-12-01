package lib

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter,status int ,data map[string]interface{}) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}