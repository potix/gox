package configure

import (
	"testing"
)

func TestLoad(t *testing.T) {
	config, err := LoadConfig("./test/test_config")
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if config.GoxHome != "/hoge" {
		t.Errorf("exp %v, act %v", "/home", config.GoxHome)
	}
	if config.ControllerAddress != "127.0.0.1" {
		t.Errorf("exp %v, act %v", "127.0.0.1", config.ControllerAddress)
	}
	if config.ControllerPort != "10000" {
		t.Errorf("exp %v, act %v", "127.0.0.1", config.ControllerPort)
	}
}

