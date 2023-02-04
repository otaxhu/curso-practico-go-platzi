package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola =)")
}

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api Endpoint =)")
}
