package main

import "net/http"

func main() {

	http.HandleFunc("/cep", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Your Cep is"))
	})

	http.ListenAndServe(":8080", nil)

}