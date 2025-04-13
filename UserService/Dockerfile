FROM golang:1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o userservice .

FROM gcr.io/distroless/static
COPY --from=builder /app/userservice /userservice
ENTRYPOINT ["/userservice"]
