package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	rpc_objects2 "rpc-and-protobuf/04-1 rpc/http-rpc/rpc_objects"
)

func main() {
	rpc.Register(new(rpc_objects2.Args))

	http.HandleFunc("/jsonRPC", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: r.Body,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":9090", nil)
}
