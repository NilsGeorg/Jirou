package main

import (
	"fmt"
	"jirou/service"
	"strconv"
)

func main() {
	fmt.Printf("Loading config...\n")
	config := service.GetConfig()

	fmt.Printf("Getting issues...\n")
	jiraResponse := service.GetJiraResponses(config)
	usedIssuesAmount := strconv.Itoa(len(jiraResponse.Issues))
	totalIssuesAmount := strconv.Itoa(jiraResponse.Total)
	fmt.Printf("The pool containes %s from a total of %s tickets!\n", usedIssuesAmount, totalIssuesAmount)

	fmt.Printf("Drawing the lucky issue...\n")
	issue := service.GetRandomIssue(jiraResponse)
	fmt.Printf("The winner is %s!\n", issue.Key)
	fmt.Printf("Summary: %s\n", issue.Fields.Summary)
	fmt.Printf("Link: %s/browse/%s\n", config.Host, issue.Key)
	fmt.Printf("Api: %s\n", issue.Self)
}
