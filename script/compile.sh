#!/bin/sh

rm -rf ./backend/proto
mkdir -p ./backend/proto

# Compile protobuf files

protoc --proto_path=./module/proto/bill-splitting \
        --go_out=./backend/proto --go_opt=paths=source_relative \
        --go-grpc_out=./backend/proto --go-grpc_opt=paths=source_relative \
        module/proto/bill-splitting/*.proto

rm -rf ./linebot/proto
mkdir -p ./linebot/proto

# Compile protobuf files

protoc --proto_path=./module/proto/bill-splitting \
        --go_out=./linebot/proto --go_opt=paths=source_relative \
        --go-grpc_out=./linebot/proto --go-grpc_opt=paths=source_relative \
        module/proto/bill-splitting/*.proto

echo "Protobuf compilation complete!"
