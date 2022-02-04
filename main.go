package main

import (
	"net/http"
	"rafaelcopat/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
