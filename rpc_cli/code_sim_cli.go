package rpc_cli

import (
	"GinAPI/config"
	"GinAPI/pb_gen"
	"google.golang.org/grpc"
	"log"
)

func initCodeSimCli() pb_gen.CodeSimClient {
	var dialOpt []grpc.DialOption
	dialOpt = append(dialOpt, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(5*1024*1024*1024*1024), grpc.MaxCallSendMsgSize(5*1024*1024*1024*1024)), grpc.WithInsecure())
	log.Printf("Dial CodeSim Addr=[%s]", config.Conf.CodeSimServiceAddr)
	conn, err := grpc.Dial(config.Conf.CodeSimServiceAddr, dialOpt...)
	if err != nil {
		panic(err)
	}
	cli := pb_gen.NewCodeSimClient(conn)
	return cli
}
