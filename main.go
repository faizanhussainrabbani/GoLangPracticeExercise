package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json: Title`
	Desc    string `json: desc`
	Content string `json: Content`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Title 1", Desc: "Desc 1", Content: "Content 1"},
	}
	fmt.Println("All articles hit")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/allarticles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
