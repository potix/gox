package router

import (
	"flag"
	"github.com/potix/plugger"
)

type Router struct {
	// ルーティングするときfilterhookを呼び出す
	filterhook func(XXXX) XXXX
}

func (g *Router)Run() {
}

func NewRouter() *Router {
	return &Router{}
}

