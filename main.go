package main

import (
	"fmt"
	"net/http"
)

func main() {

	//Handle the root url
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello world...")
	})

	//define the port listen to
	port := ":8080"

	fmt.Printf("Starting server at http://localhost%s\n", port)

	//start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
