package service

import (
    "userservice/model"
    "userservice/repository"
)

type UserService struct {
    repo repository.UserRepository // this composition not inheritance
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) AddUser(user model.User) error {
    return s.repo.Save(user)
}

func (s *UserService) GetUsers() []model.User {
    return s.repo.FindAll()
}
