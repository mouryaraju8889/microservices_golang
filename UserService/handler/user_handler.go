package handler

import (
    "encoding/json"
    "net/http"
    "userservice/model"
    "userservice/service"
)

type UserHandler struct {
    service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    _ = json.NewDecoder(r.Body).Decode(&user)
    err := h.service.AddUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    users := h.service.GetUsers()
    json.NewEncoder(w).Encode(users)
}
