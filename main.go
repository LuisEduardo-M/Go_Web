package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Web Application")

		if err != nil {
			fmt.Println(err)
		}
		bytes := fmt.Sprintf("Number of bytes written: %d", n)
		fmt.Println(bytes)
	})

	_ = http.ListenAndServe(":8080", nil)
}
