package main

import (
	"GinAPI/api"
	"GinAPI/config"
	_ "GinAPI/docs"
	"GinAPI/middlewares"
	"GinAPI/rpc_cli"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

// @title SoftwareWarehouse Web API
// @version 1.0
// @description This is a SoftwareWarehouse API server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath
func main() {
	initConfig()
	rpc_cli.InitRPCClient()
	r := gin.Default()
	registerMiddleware(r)
	api.Register(r)
	addr := fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port)
	err := r.Run(addr) // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}

func initConfig() {
	var configPath string
	var dev bool
	flag.StringVar(&configPath, "config_path", "", "配置文件路径")
	flag.BoolVar(&dev, "dev", true, "是否为测试环境")
	flag.Parse()
	env := func() config.ConfigurationEnv {
		if dev {
			return config.DevEnv
		}
		return config.PrdEnv
	}()
	config.InitConfig(configPath, env)
}

func registerMiddleware(r *gin.Engine) {
	r.NoRoute(middlewares.NotFoundHandler())
	r.NoMethod(middlewares.NotFoundHandler())
	r.Use(middlewares.Cors())
	r.Use(middlewares.Recover())
	r.Use(middlewares.ErrHandler())
}
