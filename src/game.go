package power4

import (
	"html/template"
	"log"
	"net/http"
)

// Structure du jeux
type Game struct {
	Grill      [][]int
	Player1    string
	Player2    string
	Difficulty string
}


func play(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("pages/play.html")
	if err != nil {
		log.Fatal(err)
	}
	// Prendre les valeurs du formulaire
	p1 := r.FormValue("joueur1")
	p2 := r.FormValue("joueur2")
	level := r.FormValue("difficulty")

	// Initialiser les lignes et colonnes en fonction du level choisir
	rows, cols := gridSize(level)
	
	game := Game{
		Grill:      makeGrid(rows, cols),
		Player1:    p1,
		Player2:    p2,
		Difficulty: level,
	}

	temp.Execute(w, game)
}
