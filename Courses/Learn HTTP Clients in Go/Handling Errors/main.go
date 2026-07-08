package main

import (
    "fmt"
    "net/http"
)

func fetchData(url string) (int, error) {
    res, err := http.Get(url)
    if err != nil {
        return 0, fmt.Errorf("network error: %v", err)
    }
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        return res.StatusCode, fmt.Errorf("non-OK HTTP status: %s", res.Status)
    }

    return res.StatusCode, nil
}
