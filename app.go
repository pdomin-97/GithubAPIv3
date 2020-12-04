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
	Name        string      `json:"name,omitempty"`
	FullName    string      `json:"html_url,omitempty"`
	Languague   string      `json:"languague,omitempty"`
	CreatedAt   json.Number `json:"created_at,string,omitempty"`
	PublicRepos int         `json:"public_repos,omitempty"`
	Followers   int         `json:"followers,omitempty"`
}

func main() {

	res, err := http.Get("https://api.github.com/users/DarkRaven255")
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

	var data Repos
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("failed to unmarshal:", err)
	} else {
		fmt.Println(data.FullName)
	}

}
