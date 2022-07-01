package service

import (
	"encoding/json"
	"fmt"
	"io"
	"jirou/models"
	"net/http"
	"net/url"
)

func GetJiraResponses(config models.Config) models.JiraResponse {
	request := createRequest(config)
	response := callApi(request)
	responseBody := readResponse(response)

	return transformResponse(responseBody)
}

func createRequest(config models.Config) (request *http.Request) {
	var apiPath = "/rest/api/3/search"

	request, err := http.NewRequest("GET", config.Host+apiPath, nil)
	if err != nil {
		panic(fmt.Errorf("Error while createing a new request: %w", err))
	}

	request.SetBasicAuth(config.Username, config.Token)
	request.URL.RawQuery = createQuery(config.ProjectKey, config.MaxResults).Encode()

	return
}

func createQuery(projectKey string, maxResults string) url.Values {
	var jql = fmt.Sprintf(`project = "%s" AND status = "Draft" AND resolution = "Unresolved" ORDER BY created DESC`, projectKey)
	var fields = "id,key,summary"

	var query = url.Values{}
	query.Set("fields", fields)
	query.Set("jql", jql)
	query.Set("maxResults", maxResults)

	return query
}

func callApi(request *http.Request) (response *http.Response) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(fmt.Errorf("Error while requesting jira api:%w", err))
	}

	return
}

func readResponse(response *http.Response) (responseBody []byte) {
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Errorf("Error while reading the response body: %w", err))
	}

	return
}

func transformResponse(responseBody []byte) (response models.JiraResponse) {
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		panic(fmt.Errorf("Error while transforming the response into json: %w", err))
	}

	return
}
