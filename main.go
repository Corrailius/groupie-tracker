package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func createjson(link string, name string) {
	str := readsite(link)
	os.WriteFile(name, str, 0644) // écrit un nouveau fichier, argument(nom du fichier, texte (en byte), permission)
}

func readsite(link string) []byte {
	art, err := http.Get(link) // get le lien de la page (duh)
	if err != nil {
		fmt.Println("error lmao", err)
	}
	defer art.Body.Close()
	body, err := ioutil.ReadAll(art.Body) // transforme le contenu de la page en byte
	if err != nil {
		fmt.Println("error lmao", err)
	}
	return body
}

func table(link string) { // transforme le link en dico, puis l'affiche (sur la console) note: servira pour mettre en tableau sur le site
	x := readsite(link)
	var m map[string]interface{}
	err := json.Unmarshal(x, &m) // ces deux ligne permettent de transfomer le byte (var x) en dictionnaire (stocké dans m)
	if err != nil {
		fmt.Print("lmao ERROR", err)
	}
	for _, str := range m {
		fmt.Println(str)
	}
	fmt.Println("da name:", m["name"]) // print la valeur qui a la clé name
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			return // permet de pas avoir la requette plusieurs fois (sera probablement à retirer à la fin)
		}
		table("https://groupietrackers.herokuapp.com/api/artists/1")
	})
	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe(":8080", nil)
}
