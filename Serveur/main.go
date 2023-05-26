package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"

)

func main() {
	// Gestion des routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/test", testHandler)

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur en écoute sur http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Charger le fichier home.html du dossier "Page"
	html, err := ioutil.ReadFile("./Front/Page/index.html")
	if err != nil {
		http.Error(w, "Page introuvableeee", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(html))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	// Charger le fichier test.html du dossier "Page"
	html, err := ioutil.ReadFile("../Front/Page/index2.html")
	if err != nil {
		http.Error(w, "Page introuvable", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(html))
}
