package power4

import "net/http"

func Serveur() {
	// Home route
	http.HandleFunc("/", home)

	// Play route
	http.HandleFunc("/play", play)

	// Move route
	http.HandleFunc("/move", move)

	// Rematch route
	http.HandleFunc("/rematch", rematch)

	// Load css
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Port
	http.ListenAndServe(":8080", nil)
}
