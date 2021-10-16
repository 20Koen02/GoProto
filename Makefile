go:
	protoc -I protofiles/ protofiles/*.proto \
	--go_out=./protofiles --go_opt=paths=source_relative \
	--go-grpc_out=./protofiles --go-grpc_opt=paths=source_relative