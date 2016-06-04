package router

import (
	"flag"
	"github.com/potix/plugger"
)

type Gox struct {

}

func (g *Gox)Run() {
	var config = flag.String("c", "/etc/gox.conf", "gox config file path")
}

func NewGox() *Gox {
	return &Gox
}


