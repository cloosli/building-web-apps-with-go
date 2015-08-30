package main

import (
	"net/http"

	"gopkg.in/unrolled/render.v1"
)

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Welcome, visit sub pages now."))
	})

	mux.HandleFunc("/data", func(rw http.ResponseWriter, req *http.Request) {
		r.Data(rw, http.StatusOK, []byte("some binary data here."))
	})

	mux.HandleFunc("/json", func(rw http.ResponseWriter, req *http.Request) {
		r.JSON(rw, http.StatusOK, map[string]string{"hello": "json"})
	})

	mux.HandleFunc("/html", func(rw http.ResponseWriter, req *http.Request) {
		//mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
		r.HTML(rw, http.StatusOK, "example", nil)
	})

	http.ListenAndServe(":8080", mux)
}
