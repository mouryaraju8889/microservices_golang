package repository

import "userservice/model"

type UserRepository interface {
    Save(user model.User) error
    FindAll() []model.User
}

type InMemoryUserRepo struct {
    users []model.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
    return &InMemoryUserRepo{}
}

func (r *InMemoryUserRepo) Save(user model.User) error {
    r.users = append(r.users, user)
    return nil
}

func (r *InMemoryUserRepo) FindAll() []model.User {
    return r.users
}
