package main

import (
    "fmt"
    "net/http"
    "sync"
)

func ServeIndexHtml(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/html/index.html")
}

func StartServer(wg *sync.WaitGroup, startAttempt int) {
    // Check if server failed to start and print error message if so
    if startAttempt <= 0 {
        fmt.Printf("Error: tried start the server %v, but failed\n Need to restart the server manually.", startAttempt)
    } else {
        // Decrease the WaitGroup counter when the function returns
        defer wg.Done()
        // Set up file server for static files
        fileServer := http.FileServer(http.Dir("./static"))
        http.Handle("/static/", http.StripPrefix("/static/", fileServer))
        // Serve index.html at the root path
        http.HandleFunc("/", ServeIndexHtml)
        // Start the server
        fmt.Println("http://127.0.0.1:90")
        err := http.ListenAndServe(":90", nil)
        if err != nil {
            // If server failed to start, print error message and attempt to start again
            fmt.Println(err)
            fmt.Println("Error starting the server")
            StartServer(wg, startAttempt-1)
        }
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    StartServer(&wg, 3)
    wg.Wait()
}
