package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"title"`
	Desc string  `json:"description"`
	Content string `json:"content"`
}
type Articles []Article 
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
	 Article{Title:"Test Title", Desc:"Test Desc", Content:"Hello World"},
	}
	fmt.Println("Endpoint Hit: All Article Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homePage (w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Homepage Endpoint Hit")
}
func handleRequests() {
http.HandleFunc("/", homePage)
http.HandleFunc("/articles", allArticles)
log.Fatal(http.ListenAndServe(":8081",nil))
}
func main() {
	handleRequests()
}
