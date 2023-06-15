package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("../templates/" + tmpl + ".page.tmpl")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("serveur start on : http://localhost:8080")
	http.ListenAndServe(port, nil)

}
func DataBase() {
	// Ouvrir la connexion à la base de données SQLite
	db, err := sql.Open("sqlite", "forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Lecture du fichier SQL
	sqlScript, err := ioutil.ReadFile("forum.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	// Exécution des requêtes SQL
	_, err = db.Exec(string(sqlScript))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Création des tables réussie.")
}
