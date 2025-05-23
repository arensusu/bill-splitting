
FROM golang:latest AS builder
WORKDIR /app
COPY . .
# COPY ./backend ./backend
# COPY ./script ./script

RUN apt-get update \
    && apt-get install -y protobuf-compiler \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN ./script/compile.sh

# RUN apt-get install -y gnupg wget curl unzip --no-install-recommends; \
#     wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | \
#     gpg --no-default-keyring --keyring gnupg-ring:/etc/apt/trusted.gpg.d/google.gpg --import; \
#     chmod 644 /etc/apt/trusted.gpg.d/google.gpg; \
#     echo "deb https://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list; \
#     apt-get update; \
#     apt-get install -y google-chrome-stable;

WORKDIR /app/backend
RUN go build -o main main.go
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz

FROM golang:latest
RUN mkdir /var/images
WORKDIR /app
COPY --from=builder /app/backend/main .
# COPY --from=builder /app/migrate ./migrate
# COPY backend/db/migration ./migration
# COPY backend/start.sh .
# RUN chmod +x start.sh

EXPOSE 8080 50051
# CMD ["/app/start.sh", "/app/main"]
CMD ["/app/main"]
