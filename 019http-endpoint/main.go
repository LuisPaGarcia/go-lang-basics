package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article is a json structure
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is a articles array
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func luispa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"name\":\"luispa\"}")
	fmt.Println("luispa hit")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	http.HandleFunc("/luispa", luispa)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/", homePage)

	fmt.Println("listen on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
