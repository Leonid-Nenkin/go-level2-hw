package main

import (
	"fmt"
	"time"
)

func main() {
	var switcher bool = false

	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()

	if switcher {
		go func() {
			panic("A-A-A!!!")
		}()
	}

	time.Sleep(time.Second)
}
