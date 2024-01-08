
FROM golang:latest AS builder
WORKDIR /app
COPY backend/ .
RUN go build -o main main.go

FROM golang:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY backend/Makefile .
COPY .env .

EXPOSE 8080
CMD ["/app/main"]