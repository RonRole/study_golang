package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/go/src/static")))
	http.ListenAndServe(":9090", nil)
}