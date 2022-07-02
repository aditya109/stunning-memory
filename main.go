package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		log.Print("/welcome was hit successfully !")
		fmt.Fprintf(w, "Welcome to server ! ⚡ ⛈️")
	})

	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		log.Print("/status was hit successfully !")
		fmt.Fprintf(w, "The server is running ! 🏃‍♂️⏩")
	})

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Print("/status was hit successfully !")
		fmt.Fprintf(w, "The server is running ! 🏃‍♂️⏩")
	})

	log.Print("server running successfully at http://localhost:3000 !")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("server starting failed: %v", err)
	}
}
