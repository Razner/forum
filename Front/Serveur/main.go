package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
	"golang.org/x/crypto/bcrypt"
)

type Message struct {
	Username string
	Content  string
}

type Post struct {
	Title   string
	Content string
	Image   string
}

var messages []Message
var posts []Post
var isLiked = false

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Récupérer les données du formulaire
		title := r.FormValue("title")
		content := r.FormValue("content")

		// Vérifier si les champs sont vides
		if title == "" || content == "" {
			fmt.Fprintf(w, "Veuillez remplir tous les champs")
			return
		}

		// Vérifier si la limite de caractères est respectée
		if len(content) > 300 {
			fmt.Fprintf(w, "Le contenu doit contenir au maximum 300 caractères")
			return
		}

		// Récupérer le fichier image
		file, handler, err := r.FormFile("image")
		if err != nil && err != http.ErrMissingFile {
			fmt.Fprintf(w, "Erreur lors du chargement du fichier image: %v", err)
			return
		}

		// Générer un identifiant unique pour le nom de fichier
		imageID := uuid.New().String()

		// Définir le chemin d'accès complet à l'image
		imagePath := ""
		if handler != nil {
			defer file.Close()

			imagePath = "../assets/images/" + imageID + handler.Filename

			// Sauvegarder le fichier image sur le disque
			f, err := os.Create(imagePath)
			if err != nil {
				fmt.Fprintf(w, "Erreur lors de la sauvegarde du fichier image: %v", err)
				return
			}
			defer f.Close()

			_, err = io.Copy(f, file)
			if err != nil {
				fmt.Fprintf(w, "Erreur lors de la copie du fichier image: %v", err)
				return
			}
		}

		// Créer un nouveau post avec ou sans l'image
		post := Post{Title: title, Content: content, Image: imagePath}

		// Ajouter le post à la liste des posts
		posts = append(posts, post)
	}

	// Rediriger vers la page d'accueil après la création du post
	http.Redirect(w, r, "/", http.StatusFound)
}

func likePost(w http.ResponseWriter, r *http.Request) {
	if isLiked {
		isLiked = false
		fmt.Fprintf(w, "Post unliked")
	} else {
		isLiked = true
		fmt.Fprintf(w, "Post liked")
	}
}

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

	http.Redirect(w, nil, "/", http.StatusSeeOther)
}

func general(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/general.page.tmpl"))

	data := struct {
		Posts []Post
	}{
		Posts: posts,
	}

	tmpl.Execute(w, data)
}

func login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        // Afficher la page de formulaire de connexion si la méthode HTTP n'est pas POST
        tmpl := template.Must(template.ParseFiles("../templates/login.page.tmpl"))
        tmpl.Execute(w, nil)
        return
    }

    // Récupérer les informations de connexion depuis le formulaire
    pseudo := r.FormValue("pseudo")
    password := r.FormValue("password")

    // Vérifier les informations de connexion dans la base de données
    db, err := sql.Open("sqlite3", "forum.db")
    if err != nil {
        fmt.Println(err)
        // Gérer l'erreur de manière appropriée
        return
    }
    defer db.Close()

    // Requête pour récupérer le mot de passe chiffré de l'utilisateur
    var hashedPassword string
    err = db.QueryRow("SELECT Psw FROM Users WHERE Pseudo = ?", pseudo).Scan(&hashedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("Aucun utilisateur trouvé avec ce pseudo :", pseudo)
        } else {
            fmt.Println("Erreur lors de la requête SQL :", err)
        }
        // Gérer l'erreur de manière appropriée (par exemple, afficher un message d'erreur à l'utilisateur)
        return
    }

    // Comparer le mot de passe fourni avec le mot de passe chiffré de la base de données
    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        fmt.Println("Mot de passe incorrect :", err)
        // Gérer l'erreur de manière appropriée (par exemple, afficher un message d'erreur à l'utilisateur)
        return
    }

    // Authentification réussie, rediriger vers la page d'accueil ou toute autre page souhaitée
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        // Afficher la page de formulaire d'inscription si la méthode HTTP n'est pas POST
        tmpl := template.Must(template.ParseFiles("../templates/register.page.tmpl"))
        tmpl.Execute(w, nil)
        return
    }
    pseudo := r.FormValue("pseudo")
    email := r.FormValue("email")
    password := r.FormValue("password")

    // Chiffrement du mot de passe
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Erreur lors du chiffrement du mot de passe :", err)
        // Gérer l'erreur de manière appropriée
        return
    }

    db, err := sql.Open("sqlite3", "forum.db")
    if err != nil {
        fmt.Println(err)
        // Gérer l'erreur de manière appropriée
        return
    }
    defer db.Close()

    request, err := db.Prepare("INSERT INTO Users (Pseudo, Psw, Email) VALUES(?, ?, ?)")
    if err != nil {
        fmt.Println(err)
        // Gérer l'erreur de manière appropriée
        return
    }
    _, err = request.Exec(pseudo, hashedPassword, email)
    if err != nil {
        fmt.Println(err)
        // Gérer l'erreur de manière appropriée
        return
    }

    // Redirection vers une page de confirmation
    http.Redirect(w, r, "/login", http.StatusSeeOther)
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
	http.HandleFunc("/create", createHandler)

	// Initialisez le routeur Gorilla Mux
	r := mux.NewRouter()

	// Définissez la route pour l'enregistrement d'utilisateur
	r.HandleFunc("/api/register", registerUser).Methods("POST")

	fmt.Println("Serveur démarré sur : http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type User struct {
	Pseudo string `json:"Pseudo"`
	Psw    string `json:"Psw"`
	Email  string `json:"Email"`
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
	db, err := sql.Open("sqlite3", "forum.db")
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
	sqlScript, err := ioutil.ReadFile("forum.sql")
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
