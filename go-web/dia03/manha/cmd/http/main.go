package main

import (
	"net/http"

	"github.com/bgw6/pod/cmd/http/router"
)

func main() {
	r := router.NewRouter()

	http.ListenAndServe(":8080", r.MapRoutes())
}
