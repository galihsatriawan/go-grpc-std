package std

import (
	"github.com/galihsatriawan/go-catch"
	pb "github.com/galihsatriawan/go-grpc-std/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func GrpcResponse(code int, message string, status string, data protoreflect.ProtoMessage) *pb.GRPCResponse {
	meta := &pb.Meta{
		Code:    int32(code),
		Message: message,
		Status:  status,
	}
	anyData, err := anypb.New(data)
	catch.Catch(nil, err, "Fail when try to cast to any type")
	response := pb.GRPCResponse{
		Meta: meta,
		Data: anyData,
	}

	return &response
}
func MappingResponse(src *anypb.Any, dst protoreflect.ProtoMessage) {
	err := anypb.UnmarshalTo(src, dst, proto.UnmarshalOptions{})
	catch.Catch(nil, err, "Fail when try to decode response")
}
