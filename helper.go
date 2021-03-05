package std

import (
	"github.com/galihsatriawan/go-catch"
	"github.com/galihsatriawan/go-grpc-std/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func GrpcResponse(code int, message string, status string, data protoreflect.ProtoMessage) *proto.GRPCResponse {
	meta := &proto.Meta{
		Code:    int32(code),
		Message: message,
		Status:  status,
	}
	anyData, err := anypb.New(data)
	catch.Catch(nil, err, "Fail when try to cast to any type")
	response := proto.GRPCResponse{
		Meta: meta,
		Data: anyData,
	}
	return &response
}
