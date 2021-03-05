## Preparation
- Copy protobuf-import.sh to your project directory
- Execute protobuf-import.sh
## Import protofile in destination proto
```
import "github.com/galihsatriawan/go-grpc-std/proto/response.proto";
```
## Generate Protobuf
```
protoc -I ./protobuf-import --proto_path=model --go_out=plugins=grpc:model --go_opt=paths=source_relative model/*.proto
```
