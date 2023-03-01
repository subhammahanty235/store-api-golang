package main

import (
	"fmt"
	"net/http"

	"github.com/subhammahanty235/store-api-golang/routes"
)

func main() {
	r := routes.Router()
	http.ListenAndServe(":5000", r)
	fmt.Println("Server running on port 5000")

}
