package main

import (
	"log"
	"net/http"
	"os"
)

var (
	RedirectServiceUrl string
)

func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, RedirectServiceUrl, http.StatusFound)
}

func main() {
	log.Printf("[Redirect demo service up]")
	RedirectServiceUrl = os.Getenv("REDIRECT_SERVICE_URL")
	if RedirectServiceUrl == "" {
		panic("please set REDIRECT_SERVICE_URL environment")
	}
	log.Printf("REDIRECT_SERVICE_URL env is %s", REDIRECT_SERVICE_URL)
    http.HandleFunc("/", redirect)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}