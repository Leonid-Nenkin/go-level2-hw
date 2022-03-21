package main

import "fmt"

func divide(a int, b int) int {
	return a / b
}

func main() {
	a := 1
	b := 0
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()
	divide(a, b)

}
