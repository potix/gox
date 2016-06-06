package configure

import (
	"github.com/BurntSushi/toml"
)

type immutableConfig struct {
	GoxHome           string `toml:"goxHome"`
	ControllerAddress string `toml:"controllerAddress"`
	ControllerPort    string `toml:"controllerPort"`
}

type mutableConfig struct {

}

type Config struct {
	*immutableConfig
	*mutableConfig
}

func LoadConfig(configPath string) (config *Config, err error) {
    config = new(Config)
    _, err = toml.DecodeFile(configPath, config)
    if err != nil {
        return config, err
    }
    return config, err
}
