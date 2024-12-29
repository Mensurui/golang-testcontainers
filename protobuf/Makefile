PROTO_DIRS = .
OUT_DIR = gen

proto:
	protoc -I=$(PROTO_DIRS) --go_out=$(OUT_DIR) $(PROTO_DIRS)/user.proto
