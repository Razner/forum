package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	// Définir le gestionnaire (handler) pour la racine "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Renvoyer le fichier index.html
		err := serveFile(w, r, "./forum/Front/index.html")
		if err != nil {
			http.Error(w, "Erreur de lecture du fichier HTML", http.StatusInternalServerError)
			return
		}
	})

	// Définir le gestionnaire (handler) pour "/index2"
	http.HandleFunc("/index2", func(w http.ResponseWriter, r *http.Request) {
		// Renvoyer le fichier index2.html
		err := serveFile(w, r, "./forum/Front/index2.html")
		if err != nil {
			http.Error(w, "Erreur de lecture du fichier HTML", http.StatusInternalServerError)
			return
		}
	})

	// Gérer les fichiers statiques du dossier Front
	fs := http.FileServer(http.Dir("./Front"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur en cours d'exécution sur http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
		os.Exit(1)
	}
}

func serveFile(w http.ResponseWriter, r *http.Request, filePath string) error {
	// Ouvrir le fichier
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copier le contenu du fichier dans la réponse HTTP
	_, err = io.Copy(w, file)
	if err != nil {
		return err
	}

	return nil
}
