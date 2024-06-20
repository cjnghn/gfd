package parser

import (
	"fmt"
	"net/url"
	"strings"
)

func ParseGitHubURL(gitHubURL string) (owner, repo, path string, err error) {
	parsedURL, err := url.Parse(gitHubURL)
	if err != nil {
		return "", "", "", err
	}

	parts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(parts) < 5 || parts[2] != "tree" {
		return "", "", "", fmt.Errorf("URL must be in the format: https://github.com/{owner}/{repo}/tree/{branch}/{path}")
	}

	owner = parts[0]
	repo = parts[1]
	path = strings.Join(parts[4:], "/")

	return owner, repo, path, nil
}
