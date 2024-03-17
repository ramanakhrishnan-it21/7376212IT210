package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ramanakhrishnan.K.A")
	})

	http.HandleFunc("/regno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212IT210")
	})

	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "IT")
	})

	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "red")
	})

	fmt.Printf("Server running (port=8080), routes: http://localhost:8080/hello, http://localhost:8080/name, http://localhost:8080/regno, http://localhost:8080/department, http://localhost:8080/color\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
