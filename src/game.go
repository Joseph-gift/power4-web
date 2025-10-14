package power4

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Structure du jeux
type Game struct {
	Grill      [][]int
	Player1    string
	Player2    string
	Difficulty string
	Turn       int
}

var current *Game

func play(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		// Prendre les valeurs du formulaire
		p1 := r.FormValue("joueur1")
		p2 := r.FormValue("joueur2")
		level := r.FormValue("difficulty")

		// Initialiser les lignes et colonnes en fonction du level choisir
		rows, cols := gridSize(level)

		current = &Game{
			Grill:      makeGrid(rows, cols),
			Player1:    p1,
			Player2:    p2,
			Difficulty: level,
			Turn:       1,
		}
	default:
		if current == nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	temp, err := template.ParseFiles("pages/play.html")
	if err != nil {
		log.Fatal(err)
	}

	temp.Execute(w, current)
}

func move(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/play", http.StatusSeeOther)
	}

	if current == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	colStr := r.FormValue("col")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Error(w, "invalid column", 400)
		return
	}

	placed := false
	for i := len(current.Grill) - 1; i >= 0; i-- {
		if current.Grill[i][col] == 0 {
			current.Grill[i][col] = current.Turn
			placed = true
			break
		}
	}

	if placed {
		// changer de tour
		if current.Turn == 1 {
			current.Turn = 2
		} else {
			current.Turn = 1
		}
	}

	http.Redirect(w, r, "/play", http.StatusSeeOther)
}
