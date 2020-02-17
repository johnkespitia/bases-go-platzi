package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct{
	UserId int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func main() {
	var client = &http.Client{Timeout: 10* time.Second}
	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := client.Get(url)
	if(err != nil){
		panic(err.Error()) 
	}

	var posts []Post 

	err = json.NewDecoder(resp.Body).Decode(&posts)

	if(err != nil){
		panic(err.Error()) 
	}

	fmt.Println(posts[0].Title)

}