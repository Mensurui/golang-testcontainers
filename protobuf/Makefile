PROTO_DIRS = .
OUT_DIR = gen

proto:
	protoc -I=$(PROTO_DIRS) \
	       --go_out=$(OUT_DIR) --go_opt=paths=source_relative \
	       --go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
	       $(PROTO_DIRS)/user.proto
