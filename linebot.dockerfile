FROM golang:latest
WORKDIR /app
COPY linebot/ .
COPY .env .

EXPOSE 7000
CMD ["go", "run", "main.go"]
