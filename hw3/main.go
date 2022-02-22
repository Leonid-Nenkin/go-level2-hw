package main

import (
	"fmt"
	"os"

	"github.com/Leonid-Nenkin/go-level2-hw/task3/configuration"
)

func main() {
	var cfg Configuration

	config, err := configuration.Load(cfg)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
