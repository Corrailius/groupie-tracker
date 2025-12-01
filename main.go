package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var data []Artist

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

func table(link string, n int, w http.ResponseWriter) { // args: link de l'api, la page et le serveur, l'ajoute à la variable globale data
	var x []byte
	if n == -1 {
		x = readsite(link)
	} else {
		x = readsite(link + fmt.Sprint(n))
	}
	var m Artist
	err := json.Unmarshal(x, &m) // ces deux ligne permettent de transfomer le byte (var x) en dictionnaire (stocké dans m)
	if err != nil {
		fmt.Print("lmao ERROR", err)
	}
	//fmt.Fprintf(w, "%d | %s | %d | %s\n",
	// m.ID, m.Name, m.CreationDate, m.FirstAlbum)		create a tab on the website
	data = append(data, m)
}

func printdata(w http.ResponseWriter) { // args: link de l'api, la page et le serveur, l'ajoute à la variable globale data
	for i := 0; i < len(data); i += 1 {
		fmt.Fprintf(w, "%d | %s | %d | %s\n",
			data[i].ID, data[i].Name, data[i].CreationDate, data[i].FirstAlbum) //create a tab on the website
	}
}

func tri(option string) {
	if option == "alpha" {
		for i := 1; i < len(data); i += 1 {
			for j := i; j > 0; j -= 1 {
				if data[j].Name < data[j-1].Name {
					data[j], data[j-1] = data[j-1], data[j]
				} else {
					j = -1
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			return // permet de pas avoir la requette plusieurs fois (sera probablement à retirer à la fin)
		}
		a := 0
		for a < 53 {
			if a == 0 {
				fmt.Fprintf(w, "%s | %s | %s | %s\n",
					"ID", "Name", "Creation date", "First album")
			} else {
				table("https://groupietrackers.herokuapp.com/api/artists/", a, w)
			}
			a += 1
		}
		tri("alpha")
		fmt.Print(len(data))
		printdata(w)
	})
	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe(":8080", nil)
}
