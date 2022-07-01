package models

type JiraResponse struct {
	Issues     []JiraIssue
	Expand     string
	StartAt    int
	MaxResults int
	Total      int
}
