package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
