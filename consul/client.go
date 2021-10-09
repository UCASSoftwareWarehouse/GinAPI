package consul

import (
	"GinAPI/config"
	"fmt"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"

	//_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"google.golang.org/grpc"
)

func DailWithConsulLB(serviceName string, dailOpt []grpc.DialOption) (*grpc.ClientConn, error) {
	//初始化 resolver 实例
	initResolver()
	dailOpt = append(dailOpt, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	conn, err := grpc.Dial(
		fmt.Sprintf("%s://%s/%s", "consul", config.Conf.ConsulAddr, serviceName),
		dailOpt...,
	)
	if err != nil {
		log.Println("dial err:", err)
		return conn, err
	}
	log.Println(fmt.Sprintf("gRpc consul client [%s] start success", serviceName))
	return conn, nil
}

