package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Repos type fields
type Repos struct {
	Name string `json:"name,omitempty"`
	// FullName    string      json:"html_url,omitempty"
	// Languague   string      json:"languague,omitempty"
	CreatedAt   json.Number `json:"created_at,string,omitempty"`
	PublicRepos int         `json:"public_repos,omitempty"`
	Followers   int         `json:"followers,omitempty"`
	Stars       int         `json:"stargazers_count,omitempty"`
	Watchers    int         `json:"watchers_count,omitempty"`
	Forks       int         `json:"forks,omnitempy"`
	ID          int         `json:"id,omnitempty"`
	Owner       owner       `json: "owner,omnitempty"`
	Language    string      `json: "language",omnitempty`
}

// BasicData type fields
type BasicData struct {
	Name        string      `json:"name,omitempty"`
	CreatedAt   json.Number `json:"created_at,string,omitempty"`
	PublicRepos int         `json:"public_repos,omitempty"`
	Followers   int         `json:"followers,omitempty"`
	URL         string      `json:"html_url,omitempty"`
}

type owner struct {
	Login string `json:"login"`
}

type dd []Repos

func main() {

	// res, err := http.Get("https://api.github.com/users/DarkRaven255/repos?per_page=5")
	res, err := http.Get("https://api.github.com/users/DarkRaven255/repos")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code", res.StatusCode)
	}

	var data dd
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("failed to unmarshal:", err)
	} else {
		fmt.Println(data[1])
	}

	res2, err2 := http.Get("https://api.github.com/users/DarkRaven255")
	if err2 != nil {
		log.Fatal(err2)
	}
	body2, err2 := ioutil.ReadAll(res2.Body)
	res.Body.Close()
	if err2 != nil {
		log.Fatal(err)
	}
	if res2.StatusCode != http.StatusOK {
		log.Fatal("Unexpected status code", res2.StatusCode)
	}

	var data2 BasicData
	if err2 := json.Unmarshal(body2, &data2); err2 != nil {
		fmt.Println("failed to unmarshal:", err2)
	} else {
		fmt.Println(data2)
	}

	// sb := string(body)

	// fmt.Println(sb)
}

// ZROB 2 OSOBNE FUNKCJE TRZEBA 2 GETY
// DODATKOWO OSOBNY STRUCT DLA NAME I OSOBNY DLA REPOS
