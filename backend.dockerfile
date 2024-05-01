
FROM golang:latest AS builder
WORKDIR /app
COPY backend/ .
RUN go build -o main main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

FROM golang:latest
RUN mkdir /var/images
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY backend/db/migration ./migration
COPY .env .
COPY backend/msjh.ttc .
COPY backend/start.sh .
RUN chmod +x start.sh

EXPOSE 8080
CMD ["/app/start.sh", "/app/main"]