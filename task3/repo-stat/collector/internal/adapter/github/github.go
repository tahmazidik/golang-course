package github

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"repo-stat/collector/internal/domain"

	"github.com/google/go-github/v72/github"
)

type Client struct {
	client *github.Client
}

func NewClient() *Client {
	return &Client{
		client: github.NewClient(nil),
	}
}

func parseGitHubRepoURL(repoURL string) (owner string, repo string, err error) {
	trimmedURL := strings.TrimSpace(repoURL)
	if trimmedURL == "" {
		return "", "", fmt.Errorf("repository URL cannot be empty")
	}

	parsedURL, err := url.Parse(trimmedURL)
	if err != nil {
		return "", "", fmt.Errorf("invalid repository URL: %w", err)
	}

	if parsedURL.Host != "github.com" {
		return "", "", fmt.Errorf("repository URL must point to github.com, got %s", parsedURL.Host)
	}

	pathParts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(pathParts) < 2 {
		return "", "", fmt.Errorf("repository URL must be in the format <owner>/<repo>")
	}

	return pathParts[0], pathParts[1], nil
}

func (c *Client) GetRepositoryInfo(ctx context.Context, repoURL string) (*domain.RepositoryInfo, error) {
	owner, repo, err := parseGitHubRepoURL(repoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse repository URL: %w", err)
	}

	repository, _, err := c.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to get repository info from GitHub: %w", err)
	}

	return &domain.RepositoryInfo{
		FullName:    repository.GetFullName(),
		Description: repository.GetDescription(),
		Stars:       int64(repository.GetStargazersCount()),
		Forks:       int64(repository.GetForksCount()),
		CreatedAt:   repository.GetCreatedAt().Time,
	}, nil
}
