package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		a, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			fmt.Println("error lmao", err)
		}
		defer a.Body.Close()
		body, err := ioutil.ReadAll(a.Body)
		if err != nil {
			fmt.Println("error lmao", err)
		}
		os.WriteFile("xd.json", body, 0644) // écrit un nouveau fichier, argument(nom du fichier, texte (en byte), permission)
	})
	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe(":8080", nil)
}
