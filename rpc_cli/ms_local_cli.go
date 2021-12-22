package rpc_cli

import (
	"GinAPI/config"
	"GinAPI/pb_gen"

	"google.golang.org/grpc"
)

func initMSLocalCli() pb_gen.MSLocalClient {
	var conn *grpc.ClientConn
	var dialOpt []grpc.DialOption
	dialOpt = append(dialOpt, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(5*1024*1024*1024*1024), grpc.MaxCallSendMsgSize(5*1024*1024*1024*1024)), grpc.WithInsecure())
	if config.Conf.GetEnv() == config.DevEnv {
		conn = initClientWithAddr(config.Conf.MSLocalServiceAddr, dialOpt)
		//conn = initClientWithLB(config.Conf.MSLoaclServiceName, dialOpt)
	} else {
		conn = initClientWithLB(config.Conf.MSLoaclServiceName, dialOpt)
	}
	cli := pb_gen.NewMSLocalClient(conn)
	return cli
}
