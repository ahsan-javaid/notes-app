package lib


import (
	"log"
	"net/http"
	"time"
)


func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(time.Now() ,r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
