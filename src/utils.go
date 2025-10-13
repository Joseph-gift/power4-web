package power4

import "strings"

// Fonction pour avoir les proportion de la grill
func gridSize(level string) (rows, cols int) {
	switch strings.ToLower(level) {
	case "easy":
		return 6, 7
	case "normal":
		return 6, 9
	case "hard":
		return 7, 9
	default:
		return 6, 7
	}
}

// fonction pour faire la grill 
func makeGrid(rows, cols int) [][]int {
	g := make([][]int, rows)
	for i := range g {
		g[i] = make([]int, cols)
	}
	return g
}
