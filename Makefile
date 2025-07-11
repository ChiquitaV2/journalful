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
	  --plugin=packages/grpc-client/node_modules/.bin/protoc-gen-ts_proto \
	  --ts_proto_out=./app/frontend/shared/grpc \
	  api/**/v1/*.proto

generate-sql:
	sqlc generate

build:
	go build -o bin/journalful cmd/server/main.go
