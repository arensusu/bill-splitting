FROM golang:latest
WORKDIR /app
COPY linebot/ .

EXPOSE 7000
CMD ["go", "run", "main.go"]
