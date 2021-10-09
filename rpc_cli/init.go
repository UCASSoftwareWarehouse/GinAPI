package rpc_cli

import (
	"GinAPI/consul"
	"GinAPI/pb_gen"
	"google.golang.org/grpc"
	"log"
)

var (
	CodeSimCli pb_gen.CodeSimClient
)

func InitRPCClient() {
	CodeSimCli = initCodeSimCli()
}

func initClientWithAddr(serviceAddr string, dialOpt []grpc.DialOption) *grpc.ClientConn {
	log.Printf("Dial Service with Addr=[%s]", serviceAddr)
	conn, err := grpc.Dial(serviceAddr, dialOpt...)
	if err != nil {
		panic(err)
	}
	return conn
}

func initClientWithLB(serviceName string,  dialOpt []grpc.DialOption) *grpc.ClientConn {
	conn, err := consul.DailWithConsulLB(serviceName, dialOpt)
	if err != nil {
		panic(err)
	}
	return conn
}