package main

import (
    "fmt"
    "net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
    fullURL := baseURL + "/" + id

    request, err := http.NewRequest("DELETE", fullURL, nil)
    if err != nil {
        return err
    }

    request.Header.Set("X-API-Key", apiKey)

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    if response.StatusCode > 299 {
        return fmt.Errorf(response.Status)
    }
    return nil
}
