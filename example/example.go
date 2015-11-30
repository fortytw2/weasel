package main

import (
	"net/http"

	"github.com/fortytw2/weasel"
)

func main() {
	// weasel.DSN = "localhost:3000"

	http.Handle("/weasel", weasel.Handler())
	http.ListenAndServe(":8080", nil)
}
