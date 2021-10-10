package config

import (
	"flag"
	"os"
)

var Conf *EachConfig

func InitConfig() {
	args := &args{}
	initCmdArgs(args)
	initEnvironArgs(args)
	Conf = parse(args)
}

func InitConfigWithFile(path string, env ConfigurationEnv) {
	args := &args{
		ConfigPath: path,
		Env:        env,
	}
	Conf = parse(args)
}

func initCmdArgs(args *args) {
	var configPath string
	var env string
	var port int
	flag.StringVar(&configPath, "config_path", "", "配置文件路径")
	flag.StringVar(&env, "env", "", "是否为测试测试环境，值为dev或prd")
	flag.IntVar(&port, "port", 0, "指定端口号，默认为配置文件中配置的端口号")
	flag.Parse()
	e, _ := convert2Env(env)
	if e != "" {
		args.Env = e
	}
	if port != 0 {
		args.Port = port
	}
	if configPath != "" {
		args.ConfigPath = configPath
	}
}

func initEnvironArgs(args *args) {
	env, ok := os.LookupEnv("ENV")
	if ok {
		e, _ := convert2Env(env)
		if e != "" {
			args.Env = e
		}
	}
}
