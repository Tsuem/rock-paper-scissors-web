package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func main() {
	// "/" default page
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	//parse our html file into the necessary format & send that to the browser
	//t: for template, err: checks for an error
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	//this func executes our template
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
