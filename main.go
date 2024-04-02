package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = "3000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello there\nThis is a simple web server running in Go.")
	})

	fmt.Printf("Server listening on %s.\nVisit %s on your browser\n", PORT, "http://127.0.0.1:"+PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
