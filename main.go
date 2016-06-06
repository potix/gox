package main

import (
	"os"
	"github.com/potix/gox/controller"
)

func main() {
	err := controller.Run()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
