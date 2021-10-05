package rpc_cli

import (
	"GinAPI/config"
	"GinAPI/pb_gen"
	"google.golang.org/grpc"
)

func initCodeSimCli() pb_gen.CodeSimClient {
	conn, err := grpc.Dial(config.Conf.CodeSimServiceAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	cli := pb_gen.NewCodeSimClient(conn)
	return cli
}
