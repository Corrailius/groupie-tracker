package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

// Structure pour un artiste
type Artisttest struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func mais() {
	// Sert le CSS et autres fichiers statiques
	fs := http.FileServer(http.Dir("."))
	http.Handle("/page-Style.css", fs)

	// Sert la page HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web test html.html")
	})

	// Endpoint API pour les artistes
	http.HandleFunc("/api/artists", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			http.Error(w, "Impossible de récupérer les artistes", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Erreur lors de la lecture de l'API", http.StatusInternalServerError)
			return
		}

		// Transformation JSON en slice de structs
		var data []Artisttest
		if err := json.Unmarshal(body, &data); err != nil {
			http.Error(w, "Erreur JSON", http.StatusInternalServerError)
			return
		}

		// Tri alphabétique par nom
		sort.Slice(data, func(i, j int) bool {
			return data[i].Name < data[j].Name
		})

		// Envoi des données JSON triées
		json.NewEncoder(w).Encode(data)
	})

	fmt.Println("Serveur démarré sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
