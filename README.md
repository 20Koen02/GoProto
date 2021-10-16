# Protocol Buffers and gRPC in Go

This is an updated version (as of October 2021) of [karankumarshreds/GoProto](https://github.com/karankumarshreds/GoProto)
that you can use as a template to create Go grpc clients and servers.


How it differs from karankumarshreds/GoProto:
1. Makefile with updated protoc command
2. server: Added `pb.UnimplementedMoneyTransactionServer` to server struct because it must be embedded to have forward compatible implementations
3. Deleted person protofiles because it's not used
4. Removed comments because the code speaks for itself (and if it doesn't go read the original article)


## Installing protoc and the go plugins
1. Make sure to have protoc installed. Download the protoc zip for your OS [here](https://github.com/protocolbuffers/protobuf/releases), extract and add the bin folder to your path.
2. Install the protocol compiler plugins for Go using the following commands:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```
3. Add the bin folder in your go path (`go env GOPATH`) to your path


## Compiling protofiles
1. Install make (or if you're using windows install Make for Windows)
2. run `make`


## Running
First we run the server code:
```
go run server/server.go
```

Now open a new terminal and start the client code:
```
go run client/client.go
```

The server will have the following logs:
```
Got amount  120.15
Got from  John
For  Alice
```

Thanks, again, to karankumarshreds for his awesome article and example