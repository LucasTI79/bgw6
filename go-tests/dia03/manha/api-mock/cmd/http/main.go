package main

import (
	"net/http"

	"github.com/bgw6/pod/cmd/http/router"
	"github.com/bgw6/pod/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := database.Connect(); err != nil {
		panic(err)
	}

	defer database.GetConnection().Close()

	r := router.NewRouter()

	http.ListenAndServe(":8080", r.MapRoutes())
}
