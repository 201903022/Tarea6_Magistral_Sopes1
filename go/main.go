package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo desde Go en Kubernetes!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Servidor iniciado en puerto 8080...")
	http.ListenAndServe(":8080", nil)
}
