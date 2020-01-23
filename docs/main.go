package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./_build/html/"))
	http.Handle("/", fs)

	log.Println("Hosting docs on http://localhost:3000")
	http.ListenAndServe("127.0.0.1:3000", nil)
}
