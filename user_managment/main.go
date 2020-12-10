package main

import (
	"log"
	"net/http"
	"user_managment/routing"
)

func main() {
	r := routing.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}
