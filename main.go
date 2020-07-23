package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"
)

type Post struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Content string `json:"content"`
}

var Posts []Post

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPosts(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllPosts")
	json.NewEncoder(w).Encode(Posts)
}

func returnSinglePost(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnSinglePost")
	vars := mux.Vars(r)
	key := vars["id"]
	
	for _, post := range Posts {
		if post.Id == key {
				json.NewEncoder(w).Encode(post)
		}
	}
}

func createNewPost(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println("Endpoint Hit: createNewPost")

	var post Post
	json.Unmarshal(reqBody, &post)

	Posts = append(Posts, post)
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updatePost")
	vars := mux.Vars(r)
	id := vars["id"]	
	reqBody, _ := ioutil.ReadAll(r.Body)

	var updatedPost Post
	json.Unmarshal(reqBody, &updatedPost)

	for i := range Posts {
		fmt.Println("cur post:", Posts[i])
		if Posts[i].Id == id {
			Posts[i] = updatedPost
			json.NewEncoder(w).Encode(Posts[i])
		}
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/posts", returnAllPosts)
	myRouter.HandleFunc("/post", createNewPost).Methods("POST")
	myRouter.HandleFunc("/post/{id}", updatePost).Methods("PUT")
	myRouter.HandleFunc("/post/{id}", returnSinglePost)


	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Posts = []Post{
		Post{Id: "1", Title: "Beginnings", Content: "Hello, world!"},
		Post{Id: "2", Title: "Negroni Recipe", Content: "Equal parts gin, sweet vermouth, campari."},
	}
	handleRequests()
}