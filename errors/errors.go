package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// type error interface {
//   Error() string
// }
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// Functions often return an error value, and calling code should handle errors by
	// testing whether the error equals nil.
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
