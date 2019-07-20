package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Hello World !")

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World !"))
	})
	http.ListenAndServe(":80", r)
}
