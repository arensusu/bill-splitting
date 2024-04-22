FROM golang:latest
WORKDIR /app
COPY linebot/ .
COPY .env .
RUN go build -o linebot main.go

EXPOSE 7000
CMD ["/app/linebot"]
