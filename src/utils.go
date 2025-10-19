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
		return 7, 8
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

// Fonction qui retourne la valeur de la case
func safeGet(g [][]int, r, c int) int {
	// Verifie si la positon est hors limite
	if r < 0 || r >= len(g) || c < 0 || c >= len(g[0]) {
		return 0
	}
	return g[r][c]
}

// Fonction qui compte le nombre de pion consécusifs un joueur possède
// dans une direction donnée
func countDir(g [][]int, r, c, dr, dc, player int) int {
	cnt := 0
	for {
		r += dr
		c += dc
		// Verifie si la case suivant appartient au joueur
		if safeGet(g, r, c) != player {
			break
		}
		cnt++
	}
	return cnt
}

// Fonction qui vérifie la victoire
func checkWinFrom(g [][]int, r, c, player int) bool {
	// Initaliser les directions de base
	dirs := [][2]int{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, -1},
	}
	for _, d := range dirs {
		dr, dc := d[0], d[1]
		count := 1
		count += countDir(g, r, c, dr, dc, player)
		count += countDir(g, r, c, -dr, -dc, player)
		if count >= 4 {
			return true
		}
	}
	return false
}
