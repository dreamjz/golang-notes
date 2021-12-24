package rpc_objects

import pb "rpc-and-protobuf/hello-protobuf/protobuffers"

type HelloService struct{}

func (HelloService) Hello(request *pb.User, reply *pb.User) error {
	reply.Name = "Hello: " + request.GetName()
	return nil
}
