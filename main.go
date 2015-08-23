package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	port:=os.Getenv("PORT")
	if port=="" {
		port="8080"
	}
	
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, req *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(req.FormValue("body")))
	rw.Write(markdown)
}
