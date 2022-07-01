package models

type JiraIssue struct {
	Id     string
	Key    string
	Self   string
	Expand string
	Fields JiraFields
}
