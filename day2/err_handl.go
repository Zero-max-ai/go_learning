package main

import (
	"fmt"
	"errors"
)

func errors_Re(n int) error {
	if (n < 0) {
		return errors.New("negative numbers, send positive")
	}
	return nil
}

func error_file() {
	num := -1
	err := errors_Re(num)
	if (err != nil) {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Number is positive", num)
	}
}
