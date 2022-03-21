package main

import (
	"fmt"
	"time"
)

type ErrorWithTime struct {
	error       string
	currentTime time.Time
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("Error: %s occured at %s", e.error, e.currentTime)
}

func New(text string) error {
	return &ErrorWithTime{
		error:       text,
		currentTime: time.Now(),
	}
}

func divide(a int, b int) (int, error) {
	return a / b, nil
}

func main() {
	a := 1
	b := 0

	defer func() {
		var err error
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
			err = New("custom error")
			fmt.Println(err)
		}
	}()

	_, _ = divide(a, b)

}
