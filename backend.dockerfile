
FROM golang:1.22 AS builder
WORKDIR /app
COPY backend/ .

# RUN apt-get update -y; \
#     apt-get install -y gnupg wget curl unzip --no-install-recommends; \
#     wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | \
#     gpg --no-default-keyring --keyring gnupg-ring:/etc/apt/trusted.gpg.d/google.gpg --import; \
#     chmod 644 /etc/apt/trusted.gpg.d/google.gpg; \
#     echo "deb https://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list; \
#     apt-get update -y; \
#     apt-get install -y google-chrome-stable;

RUN go build -o main main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

FROM golang:1.22
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