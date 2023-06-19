package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
)

type Message struct {
	Username string
	Content  string
}

var messages []Message

func SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	content := r.FormValue("content")

	message := Message{
		Username: username,
		Content:  content,
	}

	messages = append(messages, message)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func general(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/general.page.tmpl"))

	tmpl.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/login.page.tmpl"))

	tmpl.Execute(w, nil)
}

func register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/register.page.tmpl"))

	tmpl.Execute(w, nil)
}

func MP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/mp.page.tmpl"))

	data := struct {
		Messages []Message
	}{
		Messages: messages,
	}

	tmpl.Execute(w, data)
}

func main() {
	DataBase()

	assets := http.FileServer(http.Dir("../assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.HandleFunc("/", general)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/send", SendMessage)
	http.HandleFunc("/mp", MP)

	// Initialisez le routeur Gorilla Mux
	r := mux.NewRouter()

	// Définissez la route pour l'enregistrement d'utilisateur
	r.HandleFunc("/api/register", registerUser).Methods("POST")

	fmt.Println("Serveur démarré sur : http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

type User struct {
	ID     int    `json:"id"`
	Pseudo string `json:"pseudo"`
	Psw    string `json:"psw"`
	Email  string `json:"email"`
}

// Handler pour l'enregistrement d'utilisateur
func registerUser(w http.ResponseWriter, r *http.Request) {
	// Parsez les données JSON de la demande
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insérez les données d'utilisateur dans la base de données
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO Users (Pseudo, Psw, Email) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	_, err = statement.Exec(user.Pseudo, user.Psw, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Répondez avec un message de succès
	response := map[string]string{"message": "Enregistrement réussi"}
	json.NewEncoder(w).Encode(response)
}

func DataBase() {
	// Ouvrir la connexion à la base de données SQLite
	db, err := sql.Open("sqlite3", "forum.db")
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
