package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener la dirección IP del cliente
	clientIP := r.RemoteAddr
	// Registrar la IP del cliente en los logs del servidor
	log.Printf("Received request from %s", clientIP)
	// Responder al navegador con un simple mensaje de bienvenida
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

