FROM golang:1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o orderservice .

FROM gcr.io/distroless/static
COPY --from=builder /app/orderservice /orderservice
ENTRYPOINT ["/orderservice"]
