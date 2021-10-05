package rpc_cli

import (
	"GinAPI/config"
	"GinAPI/pb_gen"
	"context"
	"google.golang.org/grpc"
	"testing"
)

func Test_initCodeSimCli(t *testing.T) {
	config.InitConfigDefault()
	cli := initCodeSimCli()
	resp, err := cli.HelloWorld(context.Background(), &pb_gen.CodeSimHelloWorldRequest{HelloText: "hello!"})
	t.Logf("resp=[%v], err=[%v]", resp, err)
}

func Test1(t *testing.T) {
	conn, err := grpc.Dial("localhost:4401", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("err=[%v]", err)
	}
	defer conn.Close()
	cli := pb_gen.NewCodeSimClient(conn)
	res, err := cli.HelloWorld(context.Background(), &pb_gen.CodeSimHelloWorldRequest{
		HelloText:     "Hello!!!",
	})
	t.Logf("res=[%v], err=[%v]", res, err)
}
