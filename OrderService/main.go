package main

import (
    "log"
    "net/http"
    "orderservice/handler"
    "orderservice/repository"
    "orderservice/service"
)

func main() {
    repo := repository.NewInMemoryOrderRepo()
    orderService := service.NewOrderService(repo)
    orderHandler := handler.NewOrderHandler(orderService)

    http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            orderHandler.CreateOrder(w, r)
        } else if r.Method == http.MethodGet {
            orderHandler.GetOrders(w, r)
        }
    })

    log.Println("OrderService running on :8081")
    log.Fatal(http.ListenAndServe(":8081", nil))
}
