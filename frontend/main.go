package main

import (
	"fmt"
	"net/http"
)

func serveStaticFile(w http.ResponseWriter, r *http.Request) {
	// Open and read the HTML file from the directory
	http.ServeFile(w, r, "index.html")
}

func main() {
	// Registering the route and handler
	http.HandleFunc("/", serveStaticFile)

	// Starting the server in a Goroutine to handle multiple requests
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println("Server error:", err)
		}
	}()

	fmt.Println("Server started at http://localhost:8080")
	fmt.Println("Press CTRL+C to stop the server")
	select {}
}
