package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Repository struct {
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	StarsgazerCount int       `json:"stargazers_count"`
	Forks           int       `json:"forks_count"`
	CreatedAt       time.Time `json:"created_at"`
}

func main() {
	owner := flag.String("owner", "", "GitHub repository owner")
	repo := flag.String("repo", "", "GitHub repository name")
	flag.Parse()

	if *owner == "" || *repo == "" {
		log.Fatal("Both --owner and --repo flags are required")
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", *owner, *repo)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching repository information:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch repository information: %s", resp.Status)
	}

	var repository Repository

	err = json.NewDecoder(resp.Body).Decode(&repository)
	if err != nil {
		log.Fatal("Error decoding JSON response: ", err)
	}

	fmt.Printf("Repository: %s\nDescription: %s\nStars: %d\nForks: %d\nCreated At: %s\n",
		repository.Name, repository.Description, repository.StarsgazerCount, repository.Forks, repository.CreatedAt.Format("2 January 2006"))
}
