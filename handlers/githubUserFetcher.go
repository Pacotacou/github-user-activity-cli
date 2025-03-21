package handlers

import (
	"encoding/json"
	"fmt"
	"github-user-activity/models"
	"io"
	"net/http"
)

const (
	url = "https://api.github.com/users/%s/events"
)

type GithubUserFetcher struct {
	User string
	Url  string
}

func NewGithubUserFetcher(user string) (GithubUserFetcher, error) {
	guf := GithubUserFetcher{
		User: user,
		Url:  fmt.Sprintf(url, user),
	}
	return guf, nil
}

func (guf *GithubUserFetcher) FetchEvents() ([]models.Event, error) {
	req, err := http.NewRequest("GET", guf.Url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add User-Agent header as required by Github API
	req.Header.Set("User-Agent", "Github-Activity-CLI")

	//Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("user %s not found", guf.User)
		}
		return nil, fmt.Errorf("github API returned status code %d", resp.StatusCode)
	}

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse JSON response
	var events []models.Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return events, nil
}

func (guf *GithubUserFetcher) DisplayEvents() error {
	events, err := guf.FetchEvents()

	if err != nil {
		return fmt.Errorf("error fetching the events: %v", err)
	}
	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return nil
	}

	for _, event := range events {
		message := FormatEvent(event)
		if message != "" {
			fmt.Println(message)
		}
	}

	return nil
}

func FormatEvent(event models.Event) string {
	switch event.Type {
	case "PushEvent":
		commitCount := len(event.Payload.Commits)
		if commitCount > 0 {
			return fmt.Sprintf("- Pushed %d commits(s) to %s", commitCount, event.Repo.Name)
		}
	case "PullRequestEvent":
		if event.Payload.PullRequest != nil {
			return fmt.Sprintf("- %s pull request in %s: %s",
				capitalize(event.Payload.Action),
				event.Repo.Name,
				event.Payload.PullRequest.Title)
		}
	case "IssuesEvent":
		if event.Payload.Issue != nil {
			return fmt.Sprintf("- %s issue in %s: %s",
				capitalize(event.Payload.Action),
				event.Repo.Name,
				event.Payload.Issue.Title)
		}
	case "CreateEvent":
		return fmt.Sprintf("- Created %s in %s", event.Payload.RefType, event.Repo.Name)
	case "DeleteEvent":
		return fmt.Sprintf("- Deleted %s in %s", event.Payload.RefType, event.Repo.Name)
	case "WatchEvent":
		return fmt.Sprintf("- Starred %s", event.Repo.Name)
	case "ForkEvent":
		return fmt.Sprintf("- Forked %s", event.Repo.Name)
	case "IsuueCommentEvent":
		if event.Payload.Issue != nil {
			return fmt.Sprintf("- Commented on issue in %s: %s",
				event.Repo.Name,
				event.Payload.Issue.Title)
		}
	case "CommitCommentEvent":
		return fmt.Sprintf("- Commented on commit in %s", event.Repo.Name)
	case "ReleaseEvent":
		return fmt.Sprintf("- Made %s public", event.Repo.Name)
	case "MemberEvent":
		return fmt.Sprintf("- %s a collaborator to %s", capitalize(event.Payload.Action), event.Repo.Name)
	}

	return ""
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	return string(s[0]-32) + s[1:]
}
