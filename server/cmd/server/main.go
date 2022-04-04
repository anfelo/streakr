package main

import (
	"fmt"
)

// Run - Responsible for the instantiation
// and startup of our go application
func Run() error {
	fmt.Println("Starting up our application")

	return nil
}

func main() {
	fmt.Println("Streakr REST API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
