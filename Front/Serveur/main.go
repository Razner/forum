package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"github.com/google/uuid"
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
		defer file.Close()

		// Générer un identifiant unique pour le nom de fichier
		imageID := uuid.New().String()

		// Définir le chemin d'accès complet à l'image
		imagePath := "../assets/images" + imageID + handler.Filename

		// Si un fichier image a été téléchargé, le sauvegarder sur le disque
		if handler != nil {
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

		// Créer un nouveau post avec l'image
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

	http.Redirect(w, r, "/", http.StatusSeeOther)
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
	http.HandleFunc("/like", likePost)
	assets := http.FileServer(http.Dir("../assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.HandleFunc("/", general)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/send", SendMessage)
	http.HandleFunc("/mp", MP)
	http.HandleFunc("/create", createHandler)
	fmt.Println("Serveur démarré sur : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
