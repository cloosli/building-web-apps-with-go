package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

func main() {
	n := negroni.New(negroni.NewRecovery(), negroni.HandlerFunc(MyMiddleware), negroni.NewLogger(), negroni.NewStatic(http.Dir("public")))

	n.Run(":8080")
}

func MyMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there...")

	if r.URL.Query().Get("password") == "secret123" {
		next(w, r)
	} else {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
	}
	log.Println("Logging on the way back...")
}
