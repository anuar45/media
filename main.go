package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/files", http.FileServer(http.Dir("/home/anuar/media")))
	http.HandleFunc("/api", GetItemsHandler)
	http.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type MediaItem struct {
	Name  string `json:name`
	Path  string `json`
	Thumb []byte
}

func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellow World!")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "webui.html")
}
