package handler

import (
    "encoding/json"
    "net/http"
    "orderservice/model"
    "orderservice/service"
)

type OrderHandler struct {
    service *service.OrderService
}

func NewOrderHandler(s *service.OrderService) *OrderHandler {
    return &OrderHandler{service: s}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
    var order model.Order
    _ = json.NewDecoder(r.Body).Decode(&order)

    err := h.service.PlaceOrder(order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
    orders := h.service.GetOrders()
    json.NewEncoder(w).Encode(orders)
}
