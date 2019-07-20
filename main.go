package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World 3000 !")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World 3000 !"))
	})
	http.ListenAndServe(":3000", nil)
}
