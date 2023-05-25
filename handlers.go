package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}

func HanldeHome(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "api endpoint")
}