package rpc_cli

import "GinAPI/pb_gen"

var (
	CodeSimCli pb_gen.CodeSimClient
)

func InitRPCClient() {
	CodeSimCli = initCodeSimCli()
}
