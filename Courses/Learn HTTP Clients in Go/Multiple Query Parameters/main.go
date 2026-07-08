package main

import (
    "fmt"
)

func fetchTasks(baseURL, availability string) []Issue {
    var limit int
    switch availability {
    case "Low":
        limit = 1
    case "Medium":
        limit = 3
    case "High":
        limit = 5
    }

    queryParams := "?sort=estimate&" + fmt.Sprintf("limit=%d", limit)

    fullURL := baseURL + queryParams
    return getIssues(fullURL)
}
