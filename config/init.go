package config

var Conf *EachConfig

func InitConfig(configFilepath string, env ConfigurationEnv) {
	c := parse(configFilepath)
	Conf = c[env]
	Conf.Env = env
}

func setByEnviron() {

}
