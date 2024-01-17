
FROM golang:latest AS builder
WORKDIR /app
COPY backend/ .
RUN go build -o main main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

FROM golang:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY db/migration ./migration
COPY .env .
COPY start.sh .
RUN chmod +x start.sh

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]