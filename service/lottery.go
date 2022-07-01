package service

import (
	"jirou/models"
	"math/rand"
)

func GetRandomIssue(response models.JiraResponse) models.JiraIssue {
	var issueSize = len(response.Issues)
	randomIndex := rand.Intn(issueSize)

	return response.Issues[randomIndex]
}
