package rpc_cli

import (
	"GinAPI/config"
	"GinAPI/pb_gen"
	"google.golang.org/grpc"
	//"GinAPI/pb_gen"
)

func initRemoteCodeCli() pb_gen.RemoteCodeServiceClient {
	var conn *grpc.ClientConn
	var dialOpt []grpc.DialOption
	dialOpt = append(dialOpt, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(5*1024*1024*1024*1024), grpc.MaxCallSendMsgSize(5*1024*1024*1024*1024)), grpc.WithInsecure())
	if config.Conf.GetEnv() == config.DevEnv {
		conn = initClientWithAddr(config.Conf.RemoteCodeServiceAddr, dialOpt)
	} else {
		conn = initClientWithLB(config.Conf.RemoteCodeServiceName, dialOpt)
	}
	cli := pb_gen.NewRemoteCodeServiceClient(conn)
	return cli
}
