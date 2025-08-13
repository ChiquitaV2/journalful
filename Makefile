generate-proto-server:
	protoc \
      --proto_path=./api \
      --go_out=./pkg \
      --go_opt=paths=source_relative \
      --go-grpc_out=./pkg \
      --go-grpc_opt=paths=source_relative \
      api/**/v1/*.proto

generate-proto-client:
	protoc \
	  --proto_path=./api \
	  --plugin=web/node_modules/.bin/protoc-gen-ts_proto \
	  --ts_proto_opt=rpcErrorHandler=true \
      --ts_proto_out=./web/server/proto/grpc \
	  api/**/v1/*.proto

generate-sql:
	sqlc generate

build:
	go build -o bin/journalful cmd/server/main.go
