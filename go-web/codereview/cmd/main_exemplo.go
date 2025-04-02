package main

import (
	"encoding/json"
	"net/http"
)

type Usuario struct {
	Nome string
}

var usuarios []Usuario = []Usuario{
	{
		Nome: "Alexandre",
	},
}

func mainExemplo() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// encode -> transformar para json (marshal)
		// decode -> transformar num objeto (unmarshal)
		bytes, err := json.Marshal(usuarios)

		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(bytes)
	})

	http.ListenAndServe(":8080", nil)
}

// handler
// service
// repository
