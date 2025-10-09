package power4

import (
	"html/template"
	"log"
	"net/http"
)

func play(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("pages/play.html")

	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, nil)
}
