generate-proto-server:
	protoc \
      --proto_path=./api \
      --go_out=./pkg \
      --go-opt=paths=source_relative \
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

start:
	./bin/journalful

clean:
	rm -f bin/journalful

test:
	go test -v ./...

certs:
	mkdir -p testdata/certs
	openssl req -x509 -nodes -newkey rsa:2048 -keyout testdata/certs/server.key -out testdata/certs/server.crt -days 365 -subj "/CN=localhost"