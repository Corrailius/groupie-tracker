package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"locations"`
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			return // permet de pas avoir la requette plusieurs fois (sera probablement à retirer à la fin)
		}
		data = []Artist{}
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
		tri("birthday")
		printdata(w)
	})
	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe(":8080", nil)
}
