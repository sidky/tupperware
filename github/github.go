package github

type PRState int32

const (
	Open PRState = iota
	Closed
)

type PullRequest struct {
	ID          int64  `json:"id"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type GitHubClient interface {
}
