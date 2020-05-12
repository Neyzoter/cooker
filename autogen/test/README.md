# autogen
```bash
# generate go
protoc --go_out=. rpc.proto
# generate rpc service with grpc plugin
protoc --go_out=plugins=grpc:. rpc.proto
```