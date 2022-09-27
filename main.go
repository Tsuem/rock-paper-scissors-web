package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, t *http.Request) {
	html := `<strong>Hello, world</strong>`

	//this instructs the browser the kind of content it's going to receive
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprintf(w, html)
}

func main() {
	// "/" default page
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
