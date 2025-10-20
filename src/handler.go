package power4

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}
