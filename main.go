package main

import (
	"html/template"
	"net/http"
)

type Game struct {
	ID    int
	Title string
	Price float64
}

var games = []Game{
	{ID: 1, Title: "Cyberpunk 2077", Price: 59.99},
	{ID: 2, Title: "Elden Ring", Price: 49.99},
	{ID: 3, Title: "God of War", Price: 39.99},
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/purchase", purchaseHandler)
	http.HandleFunc("/game", gameHandler)

	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, games)
}

func purchaseHandler(w http.ResponseWriter, r *http.Request) {
	gameTitle := r.URL.Query().Get("title")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<p style="color: green;">Thank you for purchasing <strong>` + gameTitle + `</strong>!</p>`))
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	tmpl := template.Must(template.ParseFiles("templates/game.html"))
	tmpl.Execute(w, id)
}
