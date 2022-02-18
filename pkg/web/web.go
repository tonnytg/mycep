package web

import (
	"fmt"
	"mycep/entity/location"
	"net/http"
)

func StartWebServer() {

	http.HandleFunc("/cep", func(w http.ResponseWriter, r *http.Request) {
		// create a l
		l := location.Location{}
		// get the cep
		cep := r.URL.Query().Get("id")
		// get the l
		l.SetCEP(cep)
		// write the l
		msg := fmt.Sprintf("Success: %s", l.GetCEP())
		w.Write([]byte(msg))

	})

	http.ListenAndServe(":8080", nil)
}
