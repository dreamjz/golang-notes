package rpc_objects

import (
	"rpc-and-protobuf/04-2_protobuf/hello-protobuf/protobuffers"
)

type HelloService struct{}

func (HelloService) Hello(request *protobuffers.User, reply *protobuffers.User) error {
	reply.Name = "Hello: " + request.GetName()
	return nil
}
