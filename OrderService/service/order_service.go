package service

import (
    "fmt"
    "orderservice/client"
    "orderservice/model"
    "orderservice/repository"
)

type OrderService struct {
    repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
    return &OrderService{repo: repo}
}

func (s *OrderService) PlaceOrder(order model.Order) error {
    exists, err := client.UserExists(order.UserID)
    if err != nil {
        return err
    }

    if !exists {
        return fmt.Errorf("user not found")
    }

    return s.repo.Save(order)
}

func (s *OrderService) GetOrders() []model.Order {
    return s.repo.FindAll()
}
