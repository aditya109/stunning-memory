package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		log.Printf("/ping was hit successfully !")
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "Hi there ! ⌚ Version 2.0: 👋 My host-id is %s 🕸️", name)
	})

	log.Print("⌚ Version 2.0: server running successfully at http://localhost:3000 !")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("server starting failed: %v", err)
	}
}
