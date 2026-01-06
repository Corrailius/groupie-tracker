package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsURL string   `json:"locations"`
	DatesURL     string   `json:"concertDates"`
	RelationsURL string   `json:"relations"`
}

func main() {

	// API
	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer resp.Body.Close()

		var artists []Artist
		if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
			http.Error(w, "Erreur JSON", 500)
			return
		}

		sort.Slice(artists, func(i, j int) bool {
			return artists[i].Name < artists[j].Name
		})

		json.NewEncoder(w).Encode(artists)
	})

	// PAGE HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "page-web.html")
	})

	// CSS
	http.Handle("/page-Style.css", http.FileServer(http.Dir(".")))

	fmt.Println("Serveur sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
