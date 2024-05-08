protoc:
	protoc --go_out=./backend --go_opt=paths=source_relative --go-grpc_out=./backend --go-grpc_opt=paths=source_relative -I . proto/*.proto
	protoc --go_out=./linebot --go_opt=paths=source_relative --go-grpc_out=./linebot --go-grpc_opt=paths=source_relative -I . proto/*.proto