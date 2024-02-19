package main

import (
	handlers "bits_memory/pkg/memory/http/rest"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/memory", handlers.CardsHandler)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("../../images"))))

	port := 8080
	serverAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is listening on port: %s", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
