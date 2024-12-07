package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello, Afronautas!")
	})

	fmt.Println("Serving server on port 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}