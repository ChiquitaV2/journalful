generate-proto-client:
   protoc \
                  --proto_path=./server/proto/api \
                  --plugin=./node_modules/.bin/protoc-gen-ts_proto \
                  --ts_proto_opt=rpcErrorHandler=true \
                  --ts_proto_out=./server/proto/grpc \
                  server/proto/api/**/v1/*.proto