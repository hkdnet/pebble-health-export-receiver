package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "under construction")
	})
	http.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "method: %s", req.Method)
	})
	port := getPort()
	log.Printf("listen on port %s", port)
	http.ListenAndServe(port, nil)
}

func getPort() string {
	env := os.Getenv("PORT")
	if env == "" {
		env = "8080" // for development
	}
	return ":" + env
}
