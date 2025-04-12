package client

import (
    "encoding/json"
    "net/http"
    "orderservice/model"
)

func GetAllUsers() ([]model.User, error) {
    resp, err := http.Get("http://localhost:8080/users") // Kubernetes DNS name
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var users []model.User
    err = json.NewDecoder(resp.Body).Decode(&users)
    return users, err
}

func UserExists(userID int) (bool, error) {
    users, err := GetAllUsers()
    if err != nil {
        return false, err
    }

    for _, u := range users {
        if u.ID == userID {
            return true, nil
        }
    }

    return false, nil
}
