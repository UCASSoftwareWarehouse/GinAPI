package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type ConfigurationEnv string

const (
	DevEnv ConfigurationEnv = "dev"
	PrdEnv ConfigurationEnv = "prd"
)

type Configuration map[ConfigurationEnv]*EachConfig

type EachConfig struct {
	AppName            string `yaml:"app_name"`
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	CodeSimServiceName string `yaml:"code_sim_service_name"`
	CodeSimServiceAddr string `yaml:"code_sim_service_addr"`
	ConsulAddr string `yaml:"consul_addr"`

	Env ConfigurationEnv
}

func (c *EachConfig) GetEnv() ConfigurationEnv{
	return c.Env
}

func parse(configFilepath string) Configuration {
	if configFilepath == "" {
		envConfigPath, ok := os.LookupEnv("CONFIG_PATH")
		if !ok {
			panic("CONFIG_PATH not set, plz check your environs")
		}
		configFilepath = envConfigPath
	}
	bs, err := ioutil.ReadFile(configFilepath)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, read file failed, err=[%v]", err)
	}
	conf := make(Configuration)
	err = yaml.Unmarshal(bs, &conf)
	if err != nil {
		log.Fatalf("ConfigForEnv parse failed, unmarshal config failed, err=[%v]", err)
	}
	return conf
}
