package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func main() {
	fmt.Println("http://localhost:8080")
	http.HandleFunc("/", Home)
	http.HandleFunc("/test", Test)
	http.ListenAndServe(port, nil)

}
