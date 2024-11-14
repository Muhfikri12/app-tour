package main

import (
	"fmt"
	"net/http"
	"tour_destination/route"
)

func main() {
	db, r, _ := route.RouteInit()

	defer db.Close()

	fmt.Println("server run on port :8080")
	http.ListenAndServe(":8080", r)
}