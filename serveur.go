package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	// Home route
	http.HandleFunc("/", home)

	// Play route
	http.HandleFunc("/play", play)

	// Load css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Port
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

}
