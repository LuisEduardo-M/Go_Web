package main

import (
	"fmt"
	"github.com/LuisEduardo-M/Go_Web/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on http://localhost%s/\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
