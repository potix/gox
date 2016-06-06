package main

import (
	"github.com/potix/gox/controller"
	"os"
)

func main() {
	err := controller.Run()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
