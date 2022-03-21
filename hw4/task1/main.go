package main

import (
	"fmt"
	"time"
)

const count = 1000

func main() {
	var ch = make(chan struct{}, count)

	for i := 0; i < count; i += 1 {
		go func() {
			ch <- struct{}{}
		}()
	}
	time.Sleep(3 * time.Second)
	close(ch)

	i := 0
	for range ch {
		i += 1
	}

	fmt.Println(i)

}
