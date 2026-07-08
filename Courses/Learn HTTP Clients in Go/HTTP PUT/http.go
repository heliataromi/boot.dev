package main

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
    fullURL := baseURL + "/" + id

    jsonData, _ := json.Marshal(data)

    request, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return User{}, err
    }

    request.Header.Set("Content-Type", "application/json")
    request.Header.Set("X-API-Key", apiKey)

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        return User{}, err
    }
    defer response.Body.Close()

    var user User
    decoder := json.NewDecoder(response.Body)
    err = decoder.Decode(&user)
    if err != nil {
        return User{}, err
    }

    return user, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
    fullURL := baseURL + "/" + id

    request, err := http.NewRequest("GET", fullURL, nil)
    if err != nil {
        return User{}, err
    }

    request.Header.Set("X-API-Key", apiKey)

    response, err := http.DefaultClient.Do(request)
    if err != nil {
        return User{}, err
    }
    defer response.Body.Close()

    var user User
    decoder := json.NewDecoder(response.Body)
    err = decoder.Decode(&user)
    if err != nil {
        return User{}, err
    }

    return user, nil
}
