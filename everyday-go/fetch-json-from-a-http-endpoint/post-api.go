package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	UserID int `json:"userId"`
	ID     int `json:"id"`
	Title  string
	Body   string
}

func getPosts(apiUrl string) ([]Post, error) {
	var p []Post
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return p, err
	}
	// setup User-Agent header
	req.Header.Set("User-Agent", "posts-tutorial")
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return p, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return p, err
	}
	if err := json.Unmarshal(body, &p); err != nil {
		return p, err
	}
	return p, nil
}

func main() {
	apiURL := "https://jsonplaceholder.typicode.com/posts"
	response, err := getPosts(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range response {
		fmt.Printf("Post %d: %s\n", post.ID, post.Title)
	}
}
