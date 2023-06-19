package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Message struct {
	Username string
	Content  string
}

var messages []Message
var isLiked = false

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
	http.HandleFunc("/like", likePost)
	assets := http.FileServer(http.Dir("../assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.HandleFunc("/", general)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/send", SendMessage)
	http.HandleFunc("/mp", MP)
	fmt.Println("Serveur démarré sur : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
