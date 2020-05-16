package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/hsmtkk/heroku-go-web-app/pkg/helloworld"
)

func main() {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", helloworld.Handler)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
