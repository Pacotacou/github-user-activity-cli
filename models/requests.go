package models

import "time"

type Event struct {
	Type      string     `json:"type"`
	Repo      Repository `json:"repo"`
	Payload   Payload    `json:"payload"`
	CreatedAt time.Time  `json:"created_at"`
}

type Repository struct {
	Name string `json:"name"`
}

type Payload struct {
	Action      string       `json:"action"`
	PullRequest *PullRequest `json:"pull_request"`
	Issue       *Issue       `json:"issue"`
	Commits     []Commit     `json:"commits"`
	Ref         string       `json:"ref"`
	RefType     string       `json:"ref_type"`
}

type PullRequest struct {
	Title string `json:"title"`
}

type Issue struct {
	Title string `json:"title"`
}

type Commit struct {
	Message string `json:"message"`
}
