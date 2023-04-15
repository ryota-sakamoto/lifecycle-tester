FROM golang:1.20 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o lifecycle-tester cmd/lifecycle-tester/main.go

FROM scratch
COPY --from=builder /app/lifecycle-tester /lifecycle-tester
EXPOSE 8080
CMD ["/lifecycle-tester", "server"]
