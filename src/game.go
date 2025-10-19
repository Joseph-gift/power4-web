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
	Winner     int
	Finished   bool
	Draw       bool
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
	temp, err := template.ParseFiles("pages/play.html", "templates/result.html", "templates/board.html")
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, current)
}

func move(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}
	if current == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if current.Finished {
		http.Redirect(w, r, "/play", http.StatusAccepted)
		return
	}
	// Prendre la valeur de la colonne du button cliquer
	colStr := r.FormValue("col")
	col, err := strconv.Atoi(colStr)
	if err != nil || col < 0 || col > len(current.Grill[0]) {
		http.Error(w, "invalid column", 400)
		return
	}
	rowPlaced := -1
	// Déposer le pion dans la colonne
	for rIdx := len(current.Grill) - 1; rIdx >= 0; rIdx-- {
		if current.Grill[rIdx][col] == 0 {
			current.Grill[rIdx][col] = current.Turn
			rowPlaced = rIdx
			break
		}
	}
	// Vérifie si la colonne était pleine
	if rowPlaced == -1 {
		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}
	if checkWinFrom(current.Grill, rowPlaced, col, current.Turn) {
		current.Winner = current.Turn
		current.Finished = true
		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}
	// Vérifier s'il y a match nul
	full := true
	// Parcourir toute la grill
	for rIdx := range current.Grill {
		for cIdx := range current.Grill[rIdx] {
			// Si une case est vide full devient false
			if current.Grill[rIdx][cIdx] == 0 {
				full = false
				break
			}
		}
		if !full {
			break
		}
	}
	if full {
		current.Finished = true
		current.Draw = true
		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}
	if current.Turn == 1 {
		current.Turn = 2
	} else {
		current.Turn = 1
	}
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}
// Fonction pour recommencer la partie
func rematch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}
	if current == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	rows, cols := gridSize(current.Difficulty)
	current.Grill = makeGrid(rows, cols)
	current.Turn = 1
	current.Winner = 0
	current.Finished = false
	current.Draw = false
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}
