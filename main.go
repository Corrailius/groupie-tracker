package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     []string `json:"locations"`
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
	err := json.Unmarshal(x, &m) // ces deux ligne permettent de transfomer le byte (var x) en classe (stocké dans m)
	if err != nil {
		fmt.Print("lmao ERROR", err)
	}

	//fmt.Fprintf(w, "%d | %s | %d | %s\n",
	// m.ID, m.Name, m.CreationDate, m.FirstAlbum)		create a tab on the website
	data = append(data, m)
}

func printdata(w http.ResponseWriter) { // args: link de l'api, la page et le serveur, l'ajoute à la variable globale data
	for i := 0; i < len(data); i += 1 {
		fmt.Fprintf(w, "%d | %s | %d | %s | %d \n",
			data[i].ID, data[i].Name, data[i].CreationDate, data[i].FirstAlbum, len(data[i].Members)) //create a tab on the website
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
	if option == "birthday" {
		for i := 1; i < len(data); i += 1 {
			for j := i; j > 0; j -= 1 {
				if data[j].CreationDate < data[j-1].CreationDate {
					data[j], data[j-1] = data[j-1], data[j]
				} else {
					j = -1
				}
			}
		}
	}
	if option == "Members" {
		for i := 1; i < len(data); i += 1 {
			for j := i; j > 0; j -= 1 {
				if len(data[j].Members) > len(data[j-1].Members) {
					data[j], data[j-1] = data[j-1], data[j]
				} else {
					j = -1
				}
			}
		}
	}
	if option == "FirstAlb" {
		for i := 1; i < len(data); i += 1 {
			for j := i; j > 0; j -= 1 {
				if data[j].FirstAlbum[6:10] < data[j-1].FirstAlbum[6:10] {
					data[j], data[j-1] = data[j-1], data[j]
				} else if data[j].FirstAlbum[6:10] == data[j-1].FirstAlbum[6:10] {
					x, errr := strconv.Atoi(data[j].FirstAlbum[3:5])
					y, errr := strconv.Atoi(data[j-1].FirstAlbum[3:5])
					if errr != nil {
						// ... handle error
						panic(errr)
					}
					if x < y {
						data[j], data[j-1] = data[j-1], data[j]
					} else if x == y {
						x, errr = strconv.Atoi(data[j].FirstAlbum[0:2])
						y, errr = strconv.Atoi(data[j-1].FirstAlbum[0:2])
						if errr != nil {
							// ... handle error
							panic(errr)
						}
						if x < y {
							data[j], data[j-1] = data[j-1], data[j]
						} else {
							j = -1
						}
					} else {
						j = -1
					}
				} else {
					j = -1
				}
			}
		}
	}
}

func main() {
	// Sert le CSS et autres fichiers statiques
	fs := http.FileServer(http.Dir("."))
	http.Handle("/page-Style.css", fs)

	// Sert la page HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "page-web.html")
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
