package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("../front-end/dist"))

	//router
	http.Handle("/", fs)

	// Start the server.
	fmt.Println("Server listening on port http://localhost:8000/")
	log.Panic(
		http.ListenAndServe(":8000", nil),
	)
}