package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("🔥 DEFAULT ROUTER RUNNING 🔥")

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/auth/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("signup reached"))
	})

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
