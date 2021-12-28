package rpc_objects

type Args struct {
	M, N int
}

func (Args) Multiply(args *Args, reply *int) error {
	*reply = args.M * args.N
	return nil
}
