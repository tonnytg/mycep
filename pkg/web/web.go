package web

import (
	"encoding/json"
	"fmt"
	"log"
	"mycep/entity/location"
	"net/http"
	"os"
)

type Cep struct {
	cep string
}

type Message struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}

var (
	PORT = ":8080"
)



func StartWebServer() {

	// Health Check receive status code 200 and OK
	http.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("ok"))
	})

	http.HandleFunc("/cep", func(w http.ResponseWriter, r *http.Request) {

		l := location.NewLocation()

		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			_ = decoder.Decode(&l)
		}

		if r.Method == "GET" {
			cep := r.URL.Query().Get("id")
			if cep != "" {
				l.SetCEP(cep)
			} else {
				m := Message{Status: "error", Text: location.ErrCepEmpty.Error()}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(m)
				return
			}
		}

		// Check Valid CEP
		l.CheckCEP()

		if l.CepStatus == false {
			m := Message{Status: "error", Text: l.StatusMsg}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(m)
			return
		}

		msg := fmt.Sprintf("cep: %s", l.CEP)
		w.WriteHeader(http.StatusOK)

		m := Message{Status: "ok", Text: msg}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(m)

	})

	// Health Check receive status code 200 and OK
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("MyCEP"))
	})

	// Define PORT for webserver
	if os.Getenv("PORT") != "" {
		PORT = ":" + os.Getenv("PORT")
	}

	// Run the web server.
	log.Printf("Starting web server on port %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
