package main

import (
    "log"
    "net/http"
    "userservice/handler"
    "userservice/repository"
    "userservice/service"
)

func main() {
    repo := repository.NewInMemoryUserRepo()
    userService := service.NewUserService(repo)
    userHandler := handler.NewUserHandler(userService)

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            userHandler.CreateUser(w, r)
        } else if r.Method == http.MethodGet {
            userHandler.GetUsers(w, r)
        }
    })

    log.Println("UserService running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
