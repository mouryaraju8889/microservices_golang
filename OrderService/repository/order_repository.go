package repository

import "orderservice/model"

type OrderRepository interface {
    Save(order model.Order) error
    FindAll() []model.Order
}

type InMemoryOrderRepo struct {
    orders []model.Order
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
    return &InMemoryOrderRepo{}
}

func (r *InMemoryOrderRepo) Save(order model.Order) error {
    r.orders = append(r.orders, order)
    return nil
}

func (r *InMemoryOrderRepo) FindAll() []model.Order {
    return r.orders
}
