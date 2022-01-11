package rpc_objects

type HelloInterface interface {
	Hello(request string, reply *string) error
}
