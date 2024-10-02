package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sourabhsikarwar/go_movie_api/router"
)

func main() {
	fmt.Println("Hello, Movies API Server Starting...")
	r := router.Router()

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Printf("Server running on Port 8000")
}
