package configure

import (
	"github.com/BurntSushi/toml"
)

type immutableConfig struct {
	GoxHome           string `toml:"goxHome"`
	ControllerAddress string `toml:"controllerAddress"`
	ControllerPort    string `toml:"controllerPort"`
}

type FilterRule struct {
	Direction        string `toml:"direction"`
	SrcPluginName    string `toml:"srcPluginName"`
	DestPluginName   string `toml:"destPluginName"`
	FilterPluginName string `toml:"filterPluginName"`
}

type mutableConfig struct {
	FileterRules []FilterRule
}

type Config struct {
	*immutableConfig
	*mutableConfig
}

func Load(configPath string) (config *Config, err error) {
	config = new(Config)
	_, err = toml.DecodeFile(configPath, config)
	if err != nil {
		return config, err
	}
	return config, err
}
